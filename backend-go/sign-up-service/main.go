package main

import (
	"log"
	"net/http"
	"os"

	databaseConnection "passport-mrs-go/infrastructure/database-connection"
	router "passport-mrs-go/router"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Setup
	godotenv.Load()
	databaseConnection.InitDB()
	defer databaseConnection.CloseDB()

	// 2. Initialize Standard Go Router (Mux)
	mux := http.NewServeMux()

	// 3. Register Routes via the new separate file
	router.SignUpRoutes(mux)
	// 4. Start Server using the mux
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
