package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(email string) (string,error) {
	// Implementation for generating JWT token
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix() // Token valid for 24 hours

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(jwtSecret)
}