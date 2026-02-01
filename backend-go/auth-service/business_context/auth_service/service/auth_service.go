package business_context_service_auth_service

import (
	model "passport-mrs-go-auth-service/business_context/auth_service/model"
	repository "passport-mrs-go-auth-service/business_context/auth_service/repositories"
	infrastructure_model "passport-mrs-go-auth-service/infrastructure/model"
	infrastructure_repository "passport-mrs-go-auth-service/infrastructure/repositories"
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

	// 3. Generate JWT AccessToken
	accessToken, err2 := infrastructure_security.GenerateAccessToken(userLoginEntity.AppUserId)
	if err2 != nil {
		return infrastructure_model.BaseResponse{
			HTTPStatusCode: "500",
			Status:         "error",
			Timestamp:      time.Now().Format(time.RFC3339),
			Data:           "Failed to generate session",
		}
	}

	// 4. Generate JWT RefreshToken
	refreshToken, err3 := infrastructure_security.GenerateRefreshToken()
	if err3 != nil {
		return infrastructure_model.BaseResponse{
			HTTPStatusCode: "500",
			Status:         "error",
			Timestamp:      time.Now().Format(time.RFC3339),
			Data:           "Failed to generate session",
		}
	}

	// 5. Save RefreshToken
	err4 := infrastructure_repository.SaveRefreshToken(userLoginEntity.AppUserId, refreshToken)
	if err4 != nil {
		return infrastructure_model.BaseResponse{
			HTTPStatusCode: "500",
			Status:         "error",
			Timestamp:      time.Now().Format(time.RFC3339),
			Data:           "System failed to persist session",
		}
	}

	// 6. Success Response
	return infrastructure_model.BaseResponse{
		HTTPStatusCode: "200",
		Status:         "success",
		Timestamp:      time.Now().Format(time.RFC3339),
		Data: model.UserLoginResponseModel{
			AppUserId:    userLoginEntity.AppUserId,
			Token:        accessToken,
			RefreshToken: refreshToken,
		}, // Returning the ID as requested
	}
}
