package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/hisyntax/auth/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
	validate                         = validator.New()
)
