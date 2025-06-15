package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Debug logging
		println("🔍 CORS Debug - Origin:", origin)
		println("🔍 CORS Debug - Method:", c.Request.Method)
		println("🔍 CORS Debug - Path:", c.Request.URL.Path)

		// Allow specific origins or all origins for development
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			println("✅ CORS - Set Allow-Origin to:", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			println("✅ CORS - Set Allow-Origin to: *")
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			println("✅ CORS - Handling OPTIONS preflight request")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}
