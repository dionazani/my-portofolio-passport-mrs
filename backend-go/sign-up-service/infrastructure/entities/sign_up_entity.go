package infrastructure_entities

import (
	"time"

	"github.com/google/uuid"
)

// SignUp represents the sign_up table structure
type SignUpEntity struct {
	ID          uuid.UUID
	Fullname    string
	SignUpFrom  string // e.g., 'WEB', 'APP'
	Email       string
	MobilePhone string
	CreatedAt   time.Time
}
