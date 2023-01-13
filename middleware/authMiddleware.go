package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqquee/auth/helpers"
)

// func authentication validates the token and authorizes the user
func Authorization(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		msg := "No Authorization header provided"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Set("uid", claims.Uid)

	c.Next()

}
