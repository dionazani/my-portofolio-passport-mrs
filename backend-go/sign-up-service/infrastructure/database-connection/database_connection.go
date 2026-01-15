package infrastructure_databaseConnection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool can be accessed from other packages as db.Pool
var Pool *pgxpool.Pool

func InitDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Verify connection
	if err := Pool.Ping(context.Background()); err != nil {
		log.Fatalf("Database unreachable: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
}
