package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSONResponse(c *gin.Context, statusCode int, success bool, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"success":    success,
		"message":    message,
		"data":       data,
	})
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	JSONResponse(c, http.StatusOK, true, message, data)
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	JSONResponse(c, statusCode, false, message, nil)
}
