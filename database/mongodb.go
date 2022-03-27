package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UserCollection *mongo.Collection = OpenCollection(Client, os.Getenv("USER_COL"))
	Validate                         = validator.New()
)

func InitMongoDB() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	Mongodb := os.Getenv("MongoDB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongodb))
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Successfully connected to the mongodb server")
	return client
}

var Client *mongo.Client = InitMongoDB()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(client.Database(os.Getenv("DATABASE_NAME")).Collection(collectionName))
	return collection
}
