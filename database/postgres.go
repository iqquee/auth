package database

import (
	"fmt"
	"log"
	"os"

	"github.com/iqquee/auth/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	p_host     = os.Getenv("POSTGRES_HOST")
	p_port     = os.Getenv("POSTGRES_PORT")
	p_user     = os.Getenv("POSTGRES_USER")
	p_password = os.Getenv("POSTGRES_PWD")
	p_dbname   = os.Getenv("POSTGRES_DB_NAME")
)

var PostgresDB *gorm.DB

func InitPostgresQlDB() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", p_host, p_port, p_user, p_password, p_dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&models.User{})
	defer db.Close()
	fmt.Println("Successfully connected to postgresql server")

	return db
}
