package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "health")
	}
}
