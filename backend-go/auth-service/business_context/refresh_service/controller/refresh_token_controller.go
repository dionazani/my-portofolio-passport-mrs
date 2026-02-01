package business_context_refresh_token_service

import (
	"fmt"
	"net/http"
	infrastructure_model "passport-mrs-go-auth-service/infrastructure/model"
	infrastructure_repository "passport-mrs-go-auth-service/infrastructure/repositories"
	infrastructure_security "passport-mrs-go-auth-service/infrastructure/security"
	"time"

	"github.com/gin-gonic/gin"
)

func RefreshTokenController(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	// Validate Input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token required"})
		return
	}

	// 1. Check refresh token in DB
	appUserTokenEntity, err := infrastructure_repository.FindRefreshToken(req.RefreshToken)
	// Check if token exists and is not expired
	if err != nil || appUserTokenEntity == nil || time.Now().After(appUserTokenEntity.ExpireAt) {
		fmt.Println(err)
		fmt.Println(appUserTokenEntity)
		fmt.Println(appUserTokenEntity.ExpireAt)
		fmt.Println(appUserTokenEntity.CreatedAt)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// 2. Generate new Pair
	// Use AppUserId from the DB record to ensure we are issuing for the right person
	newAccessToken, _ := infrastructure_security.GenerateAccessToken(appUserTokenEntity.AppUserId) // Pass the role if needed
	newRefreshToken, _ := infrastructure_security.GenerateRefreshToken()

	// 3. COMPLETE THE ROTATION: Save the new token to DB
	// This will overwrite the old one because of your ON CONFLICT logic
	err = infrastructure_repository.SaveRefreshToken(appUserTokenEntity.AppUserId, newRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session"})
		return
	}

	// 4. Final Response. Inside RefreshTokenController
	c.JSON(http.StatusOK, infrastructure_model.BaseResponse{
		HTTPStatusCode: "200",
		Status:         "success",
		Timestamp:      time.Now().Format(time.RFC3339),
		Data: gin.H{
			"access_token":  newAccessToken,
			"refresh_token": newRefreshToken,
		},
	})
}
