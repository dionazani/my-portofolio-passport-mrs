package infrastructure_repository

import (
	"context"
	infrastructure_databaseConnection "passport-mrs-go-sign-up-service/infrastructure/database_connection"
	infrastructure_entitiy "passport-mrs-go-sign-up-service/infrastructure/entities" // Import from entities folder
)

func InsertSignUpAppPerson(ctx context.Context, signUpAppPersonEntity infrastructure_entitiy.SignUpAppPersonEntity) error {

	query := `INSERT INTO sign_up_app_person (sign_up_id, app_person_id) 
	          VALUES ($1, $2)`

	_, err := infrastructure_databaseConnection.Pool.Exec(ctx,
		query,
		signUpAppPersonEntity.SignUpId,
		signUpAppPersonEntity.AppPersonId)

	return err
}
