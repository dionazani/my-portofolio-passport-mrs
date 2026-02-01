package infrastructure_repositories

import (
	database_connection "passport-mrs-go-auth-service/infrastructure/database_connection"
	infrastructure_entity "passport-mrs-go-auth-service/infrastructure/entity"
	"time"

	"github.com/google/uuid" // Ensure you run: go get github.com/google/uuid
)

func SaveRefreshToken(userId string, refreshToken string) error {
	// 1. Prepare data matching your table: app_user_token
	id := uuid.New().String()
	tokenType := "refresh"
	expireAt := time.Now().Add(time.Hour * 24 * 7) // 7 Days expiration

	// 2. SQL Query using your specific columns
	// We use ON CONFLICT to ensure one user only has one active refresh session
	query := `
		INSERT INTO app_user_token (id, app_user_id, token_type, token_user, expire_at)
		VALUES (?, ?, ?, ?, ?)
		ON CONFLICT (app_user_id, token_type) 
		DO UPDATE SET 
			token_user = EXCLUDED.token_user, 
			expire_at = EXCLUDED.expire_at,
			created_at = CURRENT_TIMESTAMP
	`

	// 3. Execute via GORM Raw SQL
	return database_connection.DB.Exec(query, id, userId, tokenType, refreshToken, expireAt).Error
}

func FindRefreshToken(refreshToken string) (*infrastructure_entity.AppUserTokenEntity, error) {

	var result infrastructure_entity.AppUserTokenEntity

	// We specifically look for 'refresh' type to ensure
	// access tokens can't be used in the refresh flow.
	query := `
       SELECT id, app_user_id, token_type, token_user, expire_at, created_at
       FROM app_user_token
       WHERE token_user = ? AND token_type = 'refresh'
       LIMIT 1
    `

	// Execute the query
	db := database_connection.DB.Raw(query, refreshToken).Scan(&result)

	// 1. Check if an error occurred during the query
	if db.Error != nil {
		return nil, db.Error
	}

	// 2. Check if the token actually exists (RowsAffected will be 0 if not found)
	if db.RowsAffected == 0 {
		return nil, nil // Return nil for both to indicate "Not Found" cleanly
	}

	return &result, nil
}

func GetUserByID(userId string) (*infrastructure_entity.AppUserTokenEntity, error) {
	var appUserTokenEntity infrastructure_entity.AppUserTokenEntity

	// We use First to find the user by their UUID primary key
	// This automatically adds 'LIMIT 1' to the query
	err := database_connection.DB.Where("id = ?", userId).First(&appUserTokenEntity).Error

	if err != nil {
		return nil, err
	}

	return &appUserTokenEntity, nil
}
