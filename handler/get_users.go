package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexagonal-architecture/port/repository"
	"net/http"
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
