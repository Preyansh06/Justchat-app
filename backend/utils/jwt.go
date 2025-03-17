// backend/utils/jwt.go
package utils

import (
<<<<<<< HEAD
=======
	"errors"
>>>>>>> working-branch
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key for signing tokens (change this for security)
<<<<<<< HEAD
var jwtKey = []byte("your_secret_key")
=======
var jwtKey = []byte("root123")
>>>>>>> working-branch

// GenerateJWT creates a new JWT token for a given user email
func GenerateJWT(email string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
<<<<<<< HEAD
=======

// ValidateJWT checks if a token is valid
func ValidateJWT(tokenString string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	return claims, nil
}
>>>>>>> working-branch
