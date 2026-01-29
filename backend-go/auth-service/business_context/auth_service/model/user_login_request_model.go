package business_context_auth_service_model

import "golang.org/x/crypto/bcrypt"

type UserLoginRequestModel struct {
	Email    string `json:email`
	Password string `json:password`
}

// HashPassword takes a password and returns the bcrypt hash in a string format.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password to a hash and returns if it is valid or not.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
