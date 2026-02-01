package infrastructure_entity

import "time"

type AppUserTokenEntity struct {
	ID        string
	AppUserId string
	TokenType string
	TokenUser string
	ExpireAt  time.Time
	CreatedAt time.Time
}

func (AppUserTokenEntity) TableName() string {
	return "app_user_token"
}
