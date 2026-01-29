package database_connection

import (
	"fmt"
	_ "fmt"
	"log"
	"os"
	_ "os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	// Load .env file only if it exists (local dev).
	// In production (Docker/K8s), variables are injected directly into the OS.
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, fetching variables from system environment")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Connection Pool Settings for Production
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to setup connection pool: %v", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	log.Println("Database connection established and pool configured")
	DB = database
}
