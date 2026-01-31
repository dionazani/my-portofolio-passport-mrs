package business_context_auth_service_service

import (
	model "passport-mrs-go-auth-service/business_context/auth_service/model"
	service "passport-mrs-go-auth-service/business_context/auth_service/service"
	infra "passport-mrs-go-auth-service/infrastructure/model"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthController(c *gin.Context) {

	var req model.UserLoginRequestModel

	// 1. Bind JSON (Request handling)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, infra.BaseResponse{
			HTTPStatusCode: "400",
			Status:         "error",
			Timestamp:      time.Now().Format(time.RFC3339),
			Data:           "Invalid input format",
		})
		return
	}

	// 2. Execute Business Logic via Service
	response := service.LoginService(req)

	// 3. Deliver Response (Using your helper method for the code)
	c.JSON(response.GetIntStatusCode(), response)
}
