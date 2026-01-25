package businessContext_service_signUp

import (
	"context"
	"fmt"
	"log/slog"
	models "passport-mrs-go-sign-up-service/business-context/user-sign-up/models"
	repository_bc "passport-mrs-go-sign-up-service/business-context/user-sign-up/repository_bc"
	entity "passport-mrs-go-sign-up-service/infrastructure/entities"
	infrastructure_entity "passport-mrs-go-sign-up-service/infrastructure/entities"
	response "passport-mrs-go-sign-up-service/infrastructure/models"
	utils "passport-mrs-go-sign-up-service/utils"
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

	// signUpEntity
	signUpEntity := entity.SignUpEntity{
		ID:          signUpId,
		Fullname:    req.Fullname,
		SignUpFrom:  req.SignUpFrom,
		Email:       emailValue,
		MobilePhone: mobilePhoneValue,
		CreatedAt:   time.Now(),
	}

	// appPersonEntity
	appPersonIDGenerated := utils.GenerateUUIDV7()
	appPersonEntity := entity.AppPersonEntity{
		ID:          appPersonIDGenerated,
		Fullname:    req.Fullname,
		Email:       emailValue,
		MobilePhone: mobilePhoneValue,
		CreatedAt:   time.Now(),
	}

	// signUpApprPersonEntity
	signUpApprPersonEntity := infrastructure_entity.SignUpAppPersonEntity{
		SignUpId:    signUpEntity.ID,
		AppPersonId: appPersonEntity.ID,
	}

	// appUser
	appUserIDGenerated := utils.GenerateUUIDV7()
	appUserEntity := entity.AppUserEntity{
		ID:                     appUserIDGenerated,
		AppPersonId:            appPersonIDGenerated,
		AppUserRole:            "REG",
		AppPassword:            "MyPassword",
		MustChangePassword:     0,
		NextchangePasswordDate: time.Now().AddDate(0, 0, 120),
		IsLock:                 0,
		CreatedAt:              time.Now(),
	}

	// Persist to Database
	err = repository_bc.AddNewSignUp(ctx, signUpEntity, appPersonEntity, signUpApprPersonEntity, appUserEntity)
	if err != nil {
		slog.ErrorContext(ctx, "BusinessContext: Failed SignUp", "error", err, "ID", req.ID, "email", req.Email)

		// Return 500 Internal Server Error
		response := response.BaseResponse{
			HTTPStatusCode: "500",
			Status:         "fail",
			Timestamp:      time.Now().Format(time.RFC3339), // Standard ISO timestamp
			Data:           signUpEntity.ID,
		}

		return response, nil
	}

	// Wrap into BaseResponse (The Metadata part)
	response := response.BaseResponse{
		HTTPStatusCode: "200",
		Status:         "success",
		Timestamp:      time.Now().Format(time.RFC3339), // Standard ISO timestamp
		Data:           signUpEntity.ID,
	}

	return response, nil
}
