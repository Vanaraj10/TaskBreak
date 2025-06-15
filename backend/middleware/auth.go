package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Debug logging
		println("🔍 Auth Debug - Authorization Header:", authHeader)
		println("🔍 Auth Debug - All Headers:", c.Request.Header)
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			println("❌ Auth Error: Missing or invalid Authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing or invalid",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		println("🔍 Auth Debug - Token String:", tokenString)
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				println("❌ Auth Error: Invalid signing method:", t.Header["alg"])
				return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
			}
			jwtSecret := os.Getenv("JWT_SECRET")
			if jwtSecret == "" {
				println("❌ Auth Error: JWT_SECRET not set in environment")
			} else {
				println("✅ Auth Debug - JWT_SECRET length:", len(jwtSecret))
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			println("❌ Auth Error: Token validation failed - Error:", err)
			if token != nil {
				println("❌ Auth Error: Token Valid:", token.Valid)
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					println("❌ Auth Error: Token claims:", claims)
				}
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			println("✅ Auth Success - Email:", claims["email"])
			c.Set("email", claims["email"])
		}

		c.Next()
	}
}
