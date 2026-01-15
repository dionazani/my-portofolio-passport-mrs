package userSignUp_service

import (
	"context"
	models "passport-mrs-go/business-context/user-sign-up/models"
	entities "passport-mrs-go/infrastructure/entities"
	response "passport-mrs-go/infrastructure/models"
	repository "passport-mrs-go/infrastructure/repository"
	"time"
)

func SignUp(ctx context.Context, req models.SignUpReq) (response.BaseResponse, error) {
	// 1. Map: Request -> Entity (for Database)
	appPersonEntity := entities.AppPerson{
		ID:          req.ID,
		Fullname:    req.Fullname,
		Email:       req.Email,
		MobilePhone: req.MobilePhone,
		CreatedAt:   time.Now(),
	}

	// 2. Persist to Database
	err := repository.InsertAppPerson(ctx, appPersonEntity)
	if err != nil {
		return response.BaseResponse{}, err
	}

	// 4. Wrap into BaseResponse (The Metadata part)
	response := response.BaseResponse{
		HTTPStatusCode: "200",
		Status:         "success",
		Timestamp:      time.Now().Format(time.RFC3339), // Standard ISO timestamp
		Data:           appPersonEntity.ID,
	}

	return response, nil
}
