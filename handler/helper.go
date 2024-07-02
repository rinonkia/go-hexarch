package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func successResponse(c *gin.Context, data any) {
	c.SecureJSON(http.StatusOK, gin.H{
		"data":  data,
		"error": nil,
	})
}

func failedResponse(c *gin.Context, status int, err error) {
	c.SecureJSON(status, gin.H{
		"data":  nil,
		"error": err.Error(),
	})
}
