package infrastructure_repository

import (
	"context"
	infrastructure_databaseConnection "passport-mrs-go-sign-up-service/infrastructure/database-connection"
	infrastructure_entitiy "passport-mrs-go-sign-up-service/infrastructure/entities" // Import from entities folder
)

func InsertAppPerson(ctx context.Context, appPerson infrastructure_entitiy.AppPersonEntity) error {

	query := `INSERT INTO app_person (id, fullname, email, mobile_phone, created_at) 
	          VALUES ($1, $2, $3, $4, $5)`

	_, err := infrastructure_databaseConnection.Pool.Exec(ctx, query, appPerson.ID, appPerson.Fullname, appPerson.Email, appPerson.MobilePhone, appPerson.CreatedAt)
	return err
}
