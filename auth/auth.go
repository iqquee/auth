package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hisyntax/auth/database"
	"github.com/hisyntax/auth/helpers"
	"github.com/hisyntax/auth/interfaces"
	"github.com/hisyntax/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

var (
	Mongodb = os.Getenv("MONGO")
	MySqldb = os.Getenv("MYSQL")
)

func SignUp(c *gin.Context) {
	//check if no database was chosen
	if Mongodb == "" && MySqldb == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "you have not chosen a database yet. you must set a database you wish to use to true in the .env file to proceed",
		})
		return
	}

	if Mongodb == "yes" || Mongodb == "Yes" { //for the mongo database
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// var user database.User
		var user models.MongoUser
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

		password := interfaces.HashPassword(user.Password)
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
	} else if MySqldb == "yes" || MySqldb == "Yes" { //for the mysql database
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, userErr := interfaces.GetUserByEmail(&user, user.Email)
		if userErr != nil {
			//hash the user password
			pwdHash := interfaces.HashPassword(user.Password)
			user.Password = pwdHash

			err := interfaces.MysqlCreateUser(&user)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("this email %s has already taken by another user", user.Email),
			})
			return
		}
	}

}

func SignIn(c *gin.Context) {
	//check if no database was chosen
	if Mongodb == "" && MySqldb == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "you have not chosen a database yet. you must set a database you wish to use to true in the .env file to proceed",
		})
		return
	}

	if Mongodb == "yes" || Mongodb == "Yes" { //for the mongo database
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.Login
		var foundUser models.MongoUser

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

		passwordIsValid, msg := interfaces.VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.Email, foundUser.User_id)

		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, gin.H{
			"user": foundUser,
		})
	} else if MySqldb == "yes" || MySqldb == "Yes" { //for the mysql database
		var user models.Login
		var foundUser models.User

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		res, err := interfaces.GetUserByEmail(&foundUser, user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("the user with username %s does not exist", user.Email),
			})
			return
		}

		passwordIsValid, msg := interfaces.VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": res,
		})
	}

}
