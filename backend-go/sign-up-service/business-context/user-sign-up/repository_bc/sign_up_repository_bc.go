package businessContext_repository_signUp

import (
	"context"
	"log/slog"
	infrastructure_databaseConnection "passport-mrs-go/infrastructure/database-connection"
	infrastructure_entity "passport-mrs-go/infrastructure/entities"
)

func AddNewSignUp(ctx context.Context,
	signUpEntity infrastructure_entity.SignUpEntity,
	appPersonEntity infrastructure_entity.AppPersonEntity,
	signUpAppPersonEntity infrastructure_entity.SignUpAppPersonEntity,
	appUserEntity infrastructure_entity.AppUserEntity) error {

	// 1. Start the Transaction
	tx, err := infrastructure_databaseConnection.Pool.Begin(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to begin transaction", "error", err)
		return err
	}

	// 2. Ensure Rollback if the function exits before committing
	defer tx.Rollback(ctx)

	// --- STEP 1: Insert into sign_up ---
	querySignUp := `INSERT INTO sign_up (id, fullname, sign_up_from, email, mobile_phone, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(ctx, querySignUp, signUpEntity.ID,
		signUpEntity.Fullname,
		signUpEntity.SignUpFrom,
		signUpEntity.Email,
		signUpEntity.MobilePhone,
		signUpEntity.CreatedAt)

	if err != nil {
		slog.ErrorContext(ctx, "Repository Context: Transaction failed: InsertSignUp", "error", err)
		return err
	}

	// --- STEP 2: Insert into app_person ---
	queryAppPerson := `INSERT INTO app_person (id, fullname, email, mobile_phone, created_at) 
	          VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(ctx, queryAppPerson,
		appPersonEntity.ID,
		appPersonEntity.Fullname,
		appPersonEntity.Email,
		appPersonEntity.MobilePhone,
		appPersonEntity.CreatedAt)

	if err != nil {
		slog.ErrorContext(ctx, "Repository Context: Transaction failed: InsertAppPerson", "error", err)
		return err
	}

	// --- STEP 3: Insert into sign_up_app_person (Mapping Table) ---
	querySignUpAppPerson := `INSERT INTO sign_up_app_person (sign_up_id, app_person_id) 
	          VALUES ($1, $2)`

	_, err = tx.Exec(ctx, querySignUpAppPerson,
		signUpAppPersonEntity.SignUpId,
		signUpAppPersonEntity.AppPersonId)

	if err != nil {
		slog.ErrorContext(ctx, "Repository Context: Transaction failed: InsertSignUpAppPerson", "error", err)
		return err
	}

	// --- STEP 4: Insert into app_user ---
	queryAppUser := `INSERT INTO app_user (
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

	_, err = tx.Exec(ctx, queryAppUser, appUserEntity.ID,
		appUserEntity.AppPersonId,
		appUserEntity.AppUserRole,
		appUserEntity.AppPassword,
		appUserEntity.MustChangePassword,
		appUserEntity.NextchangePasswordDate,
		appUserEntity.IsLock,
		appUserEntity.CreatedAt)

	if err != nil {
		slog.ErrorContext(ctx, "Repository Context: Transaction failed: InsertAppUser", "error", err)
		return err
	}

	// 3. Commit the Transaction
	err = tx.Commit(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Repository Context: Failed to commit transaction", "error", err)
		return err
	}

	return nil
}
