package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv" // package used to read the .env file
	"log"
	"os"
)

// CreateConnection with postgres db
func CreateConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var (
		host     = os.Getenv("APP_DB_HOST")
		port     = os.Getenv("APP_DB_PORT")
		user     = os.Getenv("APP_DB_USERNAME")
		password = os.Getenv("APP_DB_PASSWORD")
		dbname   = os.Getenv("APP_DB_NAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
