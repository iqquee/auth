package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name    string             `json:"first_name"`
	Last_Name     string             `json:"last_name"`
	Email         string             `json:"email"`
	Phone_Number  string             `json:"phone_number"`
	Password      string             `json:"password"`
	Token         string             `json:"token"`
	Refresh_Token string             `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
}

type PublicUser struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name   string             `json:"first_name"`
	Last_Name    string             `json:"last_name"`
	Email        string             `json:"email"`
	Phone_Number string             `json:"phone_number"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
