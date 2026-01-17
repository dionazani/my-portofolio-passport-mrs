package infrastructure_entity

import (
	"time"

	"github.com/google/uuid"
)

type AppPersonEntity struct {
	ID          uuid.UUID
	Fullname    string
	Email       string
	MobilePhone string
	CreatedAt   time.Time
	UpdateAt    time.Time
}
