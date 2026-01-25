package route

import (
	"net/http"
	"os"
	businessContext_controller "passport-mrs-go-sign-up-service/business-context/user-sign-up/controllers"
)

// RegisterRoutes sets up all the endpoints for this business context
func SignUpRoutes(mux *http.ServeMux) {
	// Fetch paths from .env
	signUpPath := os.Getenv("SIGN_UP_ENDPOINT")
	if signUpPath == "" {
		signUpPath = "/api/v1/sign-up" // Fallback
	}

	// Register handlers
	mux.HandleFunc(signUpPath, businessContext_controller.SignUpHandler)

	// You can easily add more routes here in the future
	// mux.HandleFunc("/api/v1/login", LoginHandler)
}
