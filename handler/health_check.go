package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
