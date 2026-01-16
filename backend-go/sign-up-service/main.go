package main

import (
	"log"
	"net/http"
	"os"

	databaseConnection "passport-mrs-go/infrastructure/database-connection"
	infrastructure_logger "passport-mrs-go/infrastructure/logger"
	router "passport-mrs-go/infrastructure/router"
	middleware "passport-mrs-go/middleware"

	"github.com/joho/godotenv"
)

func main() {

	// setup logger
	infrastructure_logger.InitLogger()

	// 1. Setup
	godotenv.Load()
	databaseConnection.InitDB()
	defer databaseConnection.CloseDB()

	// 2. Initialize Standard Go Router (Mux)
	mux := http.NewServeMux()

	// 3. Register Routes via the new separate file
	router.SignUpRoutes(mux)

	// --- MIDDLEWARE WRAPPING ---
	// Wrap the mux with the RequestLogger middleware
	wrappedMux := middleware.RequestLogger(mux)

	// 4. Start Server using the mux
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8089"
	}

	log.Printf("🚀 Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, wrappedMux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
