package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(email string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		println("❌ JWT Generation Error: JWT_SECRET not set")
		return "", fmt.Errorf("JWT_SECRET not set")
	}
	println("✅ JWT Generation - Secret length:", len(jwtSecret))
	println("✅ JWT Generation - Email:", email)

	// Implementation for generating JWT token
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix() // Token valid for 72 hours
	claims["iat"] = time.Now().Unix()                     // Issued at time

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		println("❌ JWT Generation Error:", err)
		return "", err
	}

	println("✅ JWT Generation - Token created successfully")
	return tokenString, nil
}
