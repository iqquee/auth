package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hisyntax/auth/database"
	"github.com/hisyntax/auth/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "Incorrect Password Please try again"
		check = false
	}

	return check, msg
}

const (
	Mongodb    = "Mongodb"
	Postgresdb = "Postgresdb"
)

func SignUp(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user database.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.Validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filter := bson.D{{Key: "email", Value: user.Email}}
	count, err := database.UserCollection.CountDocuments(ctx, filter)
	if err != nil {
		log.Panic(err)
		msg := "Error occured while checking for the Email"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	if count > 0 {
		msg := "this email already exists"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	password := HashPassword(user.Password)
	user.Password = password

	user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	token, refreshToken, _ := helpers.GenerateAllTokens(user.Email, user.User_id)
	user.Token = token
	user.Refresh_Token = refreshToken

	insertNum, insertErr := database.UserCollection.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": insertErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": insertNum,
	})

}

func SignIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user database.Login
	var foundUser database.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filter := bson.D{{Key: "email", Value: user.Email}}
	err := database.UserCollection.FindOne(ctx, filter).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.Email, foundUser.User_id)

	helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	c.JSON(http.StatusOK, gin.H{
		"user": foundUser,
	})
}
