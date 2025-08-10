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
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file:", err)
	}

	// Get environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")

	// Build DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	// Connect to MySQL
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// Test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database ping failed:", err)
	}

	log.Println("✅ Connected to database!")
}
