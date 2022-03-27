package models

import (
	"time"

	// valid "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUser struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	User_id       string             `json:"user_id"`
	First_Name    string             `json:"first_name"`
	Last_Name     string             `json:"last_name"`
	Email         string             `json:"email"`
	Phone_Number  int                `json:"phone_number"`
	Password      string             `json:"password"`
	Token         string             `json:"token"`
	Refresh_Token string             `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
}

type User struct {
	gorm.Model
	First_Name    string `json:"first_name"`
	Last_Name     string `json:"last_name"`
	Email         string `json:"email"`
	Phone_Number  int    `json:"phone_number"`
	Password      string `json:"password"`
	Token         string `json:"token"`
	Refresh_Token string `json:"refresh_token"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
