package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a password and returns the bcrypt hash in a string format.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
