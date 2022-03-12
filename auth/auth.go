package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hisyntax/auth/helpers"
	"github.com/hisyntax/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.Validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filter := bson.D{{Key: "email", Value: user.Email}}
	count, err := utils.UserCollection.CountDocuments(ctx, filter)
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

	password := utils.HashPassword(user.Password)
	user.Password = password

	user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	token, refreshToken, _ := helpers.GenerateAllTokens(user.Email, user.User_id)
	user.Token = token
	user.Refresh_Token = refreshToken

	insertNum, insertErr := utils.UserCollection.InsertOne(ctx, user)
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

	var user Login
	var foundUser User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filter := bson.D{{Key: "email", Value: user.Email}}
	err := utils.UserCollection.FindOne(ctx, filter).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	passwordIsValid, msg := utils.VerifyPassword(foundUser.Password, user.Password)
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
