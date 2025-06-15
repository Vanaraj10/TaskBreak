package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse sends a standardized error response
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"success": false,
		"error":   message,
	})
}

// SuccessResponse sends a standardized success response
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// SuccessMessageResponse sends a success response with a message
func SuccessMessageResponse(c *gin.Context, message string, data interface{}) {
	response := gin.H{
		"success": true,
		"message": message,
	}

	if data != nil {
		response["data"] = data
	}

	c.JSON(http.StatusOK, response)
}
