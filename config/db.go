package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {

	// // Load environment variables from .env file (DEVELOPMENT ONLY)
	// err := godotenv.Load(".env")

	// Load environment variables from runtime (PRODUCTION)
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// Get database credentials from environment variables
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Open the database connection
	db, err := sql.Open("mysql", dsn)

	// // Development database
	// db, err := sql.Open("mysql", "root:@/naturemate")

	if err != nil {
		return nil, err
	}

	// Test connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection successful")

	return db, nil
}
