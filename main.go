package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file missing")
	}
}

func main() {
	r := gin.Default()

	auths := r.Group("/auth")
	{
		auths.POST("/signup")
		auths.POST("/signin")
	}
}
