// backend/utils/jwt.go
package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key for signing tokens (change this for security)
var jwtKey = []byte("your_secret_key")

// GenerateJWT creates a new JWT token for a given user email
func GenerateJWT(email string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
