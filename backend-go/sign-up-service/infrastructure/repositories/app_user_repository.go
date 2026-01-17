package infrastructure_repository

import (
	"context"
	infrastructure_databaseConnection "passport-mrs-go/infrastructure/database-connection"
	entity "passport-mrs-go/infrastructure/entities"
)

func InsertAppUser(ctx context.Context, appUserEntity entity.AppUserEntity) error {

	query := `INSERT INTO app_user (
				id, 
				app_person_id, 
				app_user_role, 
				app_password, 
				must_change_password, 
				next_change_password_date, 
				is_locked, 
				created_at
			  ) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := infrastructure_databaseConnection.Pool.Exec(
		ctx,
		query,
		appUserEntity.ID,
		appUserEntity.AppPersonId,
		appUserEntity.AppUserRole,
		appUserEntity.AppPassword,
		appUserEntity.MustChangePassword,
		appUserEntity.NextchangePasswordDate,
		appUserEntity.IsLock,
		appUserEntity.CreatedAt,
	)

	return err
}
