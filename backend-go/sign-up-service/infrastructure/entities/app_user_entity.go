package infrastructure_entity

import (
	"time"

	"github.com/google/uuid"
)

type AppUserEntity struct {
	ID                     uuid.UUID
	AppPersonId            uuid.UUID
	AppUserRole            string
	AppPassword            string
	MustChangePassword     int
	NextchangePasswordDate time.Time
	IsLock                 int
	CreatedAt              time.Time
}
