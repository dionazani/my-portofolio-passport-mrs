package business_context_auth_service_service

import (
	"net/http"
	model "passport-mrs-go-auth-service/business_context/auth_service/model"
	repository "passport-mrs-go-auth-service/business_context/auth_service/repositories"
	model_infrastructure "passport-mrs-go-auth-service/infrastructure/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthController(c *gin.Context) {

	var req model.UserLoginRequestModel
	// 1. Validate JSON Input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	userLoginEntity, err := repository.GetUserLogin(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	var isPasswordMatch bool = model.CheckPasswordHash(req.Password, userLoginEntity.AppPassword)
	if !isPasswordMatch {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	response := model_infrastructure.BaseResponse{
		HTTPStatusCode: "200",
		Status:         "success",
		Timestamp:      time.Now().Format(time.RFC3339),
		Data:           userLoginEntity.AppUserId,
	}

	httpStatusCode, err := strconv.Atoi(response.HTTPStatusCode) // Atoi returns (int, error)
	if err != nil {
		// Fallback to 200 if conversion fails, or send 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(httpStatusCode, response)
}
