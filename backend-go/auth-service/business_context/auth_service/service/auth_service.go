package business_context_auth_service_service

import (
	model "passport-mrs-go-auth-service/business_context/auth_service/model"
	repository "passport-mrs-go-auth-service/business_context/auth_service/repositories"
	infrastructure_model "passport-mrs-go-auth-service/infrastructure/model"
	infrastructure_security "passport-mrs-go-auth-service/infrastructure/security"

	"time"
)

func LoginService(req model.UserLoginRequestModel) infrastructure_model.BaseResponse {
	// 1. Fetch User
	userLoginEntity, err := repository.GetUserLogin(req.Email)
	if err != nil {
		return infrastructure_model.BaseResponse{
			HTTPStatusCode: "401",
			Status:         "unauthorized",
			Timestamp:      time.Now().Format(time.RFC3339),
			Data:           "Invalid email or password",
		}
	}

	// 2. Verify Password
	isPasswordMatch := model.CheckPasswordHash(req.Password, userLoginEntity.AppPassword)
	if !isPasswordMatch {
		return infrastructure_model.BaseResponse{
			HTTPStatusCode: "401",
			Status:         "unauthorized",
			Timestamp:      time.Now().Format(time.RFC3339),
			Data:           "Invalid email or password",
		}
	}

	// 3. Generate JWT Token
	token, err2 := infrastructure_security.GenerateToken(userLoginEntity.AppUserId)
	if err2 != nil {
		return infrastructure_model.BaseResponse{
			HTTPStatusCode: "500",
			Status:         "error",
			Timestamp:      time.Now().Format(time.RFC3339),
			Data:           "Failed to generate session",
		}
	}

	// 4. Success Response
	return infrastructure_model.BaseResponse{
		HTTPStatusCode: "200",
		Status:         "success",
		Timestamp:      time.Now().Format(time.RFC3339),
		Data: model.UserLoginResponseModel{
			AppUserId: userLoginEntity.AppUserId,
			Token:     token,
		}, // Returning the ID as requested
	}
}
