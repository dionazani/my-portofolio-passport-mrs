package infrastructure_security

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken creates a new JWT for a specific User ID
func GenerateAccessToken(userId string) (string, error) {

	// secret key should in .env
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "Anson-phantom-morgue-improve-acres-urged-overcome-clients-iast-britney-kicker-sol-acids-weaving-post"
	}

	// Create the claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
