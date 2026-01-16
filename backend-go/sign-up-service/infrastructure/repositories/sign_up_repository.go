package infrastructure_repository

import (
	"context"
	infrastructure_databaseConnection "passport-mrs-go/infrastructure/database-connection"
	entitiy "passport-mrs-go/infrastructure/entities" // Import from entities folder
)

func InsertSignUp(ctx context.Context, signUpEntity entitiy.SignUpEntity) error {

	query := `INSERT INTO sign_up (id, fullname, sign_up_from, email, mobile_phone, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := infrastructure_databaseConnection.Pool.Exec(ctx, query, signUpEntity.ID, signUpEntity.Fullname, signUpEntity.SignUpFrom, signUpEntity.Email, signUpEntity.MobilePhone, signUpEntity.CreatedAt)
	return err
}
