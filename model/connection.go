package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Db is instance of model
var Db *sql.DB

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Db, err = sql.Open("mysql", os.Getenv("MYSQL_CRED"))
	if err != nil {
		log.Fatal(err)
	}
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfuly connected...")
}
