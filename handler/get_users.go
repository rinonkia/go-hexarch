package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/port/repository"
)

func GetUsers(repo repository.UserRepository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		us := repo.GetAll()
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"users":   us,
		})
	}
}
