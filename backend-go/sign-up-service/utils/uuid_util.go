package utils

import (
	"log"

	"github.com/google/uuid"
)

func GenerateUUIDV7() uuid.UUID {

	// Generate a new Version 7 UUID
	u7, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	return u7
}
