package db

import (
	"fmt"
	"database/sql"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DB() (*sql.DB, error) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
}
