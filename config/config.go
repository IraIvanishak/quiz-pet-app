package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var ConnStr string

func init() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	ConnStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
	}
	DB = db
}
