package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hisyntax/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPublicUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundUser []PublicUser

	userFilter := bson.D{{}}
	userCol, err := utils.UserCollection.Find(ctx, userFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := userCol.All(ctx, &foundUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer userCol.Close(ctx)

	c.JSON(http.StatusOK, gin.H{
		"users": foundUser,
	})
}

func GetPublicUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundUser PublicUser
	userEmail := c.Request.Header.Get("email")
	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user email needs to be provided in the header",
		})
		return
	}

	userFilter := bson.D{{Key: "email", Value: userEmail}}
	userCol := utils.UserCollection.FindOne(ctx, userFilter)
	if err := userCol.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := userCol.Decode(&foundUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": foundUser,
	})
}
