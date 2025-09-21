package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("✅ Connected to MySQL database")
}
