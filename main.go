package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "giftOF463"
	dbname   = "mydb"
)

func main() {
	// database.SetupDB()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)

	err = db.Ping()
	checkError(err)

	defer db.Close()
	fmt.Println("Successfully connected!")
}
