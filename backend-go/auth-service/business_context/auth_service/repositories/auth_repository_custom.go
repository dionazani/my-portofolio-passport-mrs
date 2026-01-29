package business_context_auth_service_repositories

import (
	entity "passport-mrs-go-auth-service/business_context/auth_service/entities"
	config "passport-mrs-go-auth-service/infrastructure/database_connection"
)

func GetUserLogin(email string) (*entity.UserLoginAuthEntity, error) {

	var userLogin entity.UserLoginAuthEntity

	// The query uses ? as a placeholder for parameters
	query := `select app_person.email, 
       					app_user.app_password, 
       					app_user.id as app_user_id, 
       					app_user.is_locked 
				from app_person inner join app_user on app_person.id = app_user.app_person_id 
				where app_person.email = ?`

	// config.DB is your *gorm.DB instance
	err := config.DB.Raw(query, email).Scan(&userLogin).Error

	if err != nil {
		return nil, err
	}

	return &userLogin, nil
}
