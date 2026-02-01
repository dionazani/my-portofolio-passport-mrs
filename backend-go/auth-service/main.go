package main

import (
	"os"
	controller_auth "passport-mrs-go-auth-service/business_context/auth_service/controller"
	controller_refresh "passport-mrs-go-auth-service/business_context/refresh_service/controller"
	database_connection "passport-mrs-go-auth-service/infrastructure/database_connection"

	"github.com/gin-gonic/gin"
)

func main() {
	database_connection.ConnectToDatabase()

	// Use Gin ReleaseMode in production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// Grouping routes is a best practice for Auth Services
	auth := r.Group("/api/v1/auth")
	{
		// Map the POST request to your AuthController
		auth.POST("/login", controller_auth.AuthController)

		// Map the POST request to your AuthController
		auth.POST("/refresh", controller_refresh.RefreshTokenController)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090" // Default fallback
	}

	r.Run(":" + port)
}
