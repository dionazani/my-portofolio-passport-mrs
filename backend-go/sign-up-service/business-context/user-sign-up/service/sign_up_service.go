package userSignUp_service

import (
	"context"
	"fmt"
	"log/slog"
	models "passport-mrs-go/business-context/user-sign-up/models"
	entities "passport-mrs-go/infrastructure/entities"
	response "passport-mrs-go/infrastructure/models"
	repository "passport-mrs-go/infrastructure/repositories"
	"time"

	"github.com/google/uuid"
)

func SignUp(ctx context.Context, req models.SignUpReqModel) (response.BaseResponse, error) {

	// Simple, shared logging call
	slog.InfoContext(ctx, "Business Logic: Sign Up", "ID", req.ID, "email", req.Email)

	var emailValue string
	if req.Email != nil {
		emailValue = *req.Email
	} else {
		emailValue = ""
	}

	var mobilePhoneValue string
	if req.MobilePhone != nil {
		mobilePhoneValue = *req.MobilePhone
	} else {
		mobilePhoneValue = ""
	}

	// Convert the string to a uuid.UUID
	signUpId, err := uuid.Parse(req.ID)
	if err != nil {
		// Handle the error (e.g., log it, return from function, etc.)
		fmt.Printf("Error parsing UUID: %v\n", err)
		return response.BaseResponse{}, err
	}

	// 1. Map: Request -> Entity (for Database)
	signUpEntity := entities.SignUpEntity{
		ID:          signUpId,
		Fullname:    req.Fullname,
		Email:       emailValue,
		MobilePhone: mobilePhoneValue,
		CreatedAt:   time.Now(),
	}

	// 2. Persist to Database
	err = repository.InsertSignUp(ctx, signUpEntity)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to save user sign-up to database", "error", err, "ID", req.ID, "email", req.Email)
		return response.BaseResponse{}, err
	}

	// 4. Wrap into BaseResponse (The Metadata part)
	response := response.BaseResponse{
		HTTPStatusCode: "200",
		Status:         "success",
		Timestamp:      time.Now().Format(time.RFC3339), // Standard ISO timestamp
		Data:           signUpEntity.ID,
	}

	return response, nil
}
