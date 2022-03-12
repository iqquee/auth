package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type PublicUser struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name   string             `json:"first_name"`
	Last_Name    string             `json:"last_name"`
	Email        string             `json:"email"`
	Phone_Number string             `json:"phone_number"`
}
