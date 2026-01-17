package infrastructure_entity

import (
	"github.com/google/uuid"
)

type SignUpAppPersonEntity struct {
	SignUpId    uuid.UUID
	AppPersonId uuid.UUID
}
