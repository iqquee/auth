package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserPsql struct {
	gorm.Model
	First_Name    string `json:"first_name"`
	Last_Name     string `json:"last_name"`
	Email         string `json:"email"`
	Phone_Number  int    `json:"phone_number"`
	Password      string `json:"password"`
	Token         string `json:"token"`
	Refresh_Token string `json:"refresh_token"`
}

var (
	p_host     = os.Getenv("POSTGRES_HOST")
	p_port     = os.Getenv("POSTGRES_PORT")
	p_user     = os.Getenv("POSTGRES_USER")
	p_password = os.Getenv("POSTGRES_PWD")
	p_dbname   = os.Getenv("POSTGRES_DB_NAME")
)

func InitPostgresQlDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", p_host, p_port, p_user, p_password, p_dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	CheckErr(err)

	db.AutoMigrate(&UserPsql{})
	defer db.Close()
	fmt.Println("Successfully connected to postgresql server")

	return db
}

func CheckErr(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
