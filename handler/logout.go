package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}

}
