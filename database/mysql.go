package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func BuildDbConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PWD"),
		DBName:   os.Getenv("MYSQL_DB_NAME"),
	}
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
