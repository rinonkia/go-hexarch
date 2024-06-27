package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexagonal-architecture/core/model"
	"github.com/rinonkia/go-hexagonal-architecture/port/repository"
)

func Signup(repo repository.UserRepository) func(ctx *gin.Context) {
	u := &model.User{
		Name:     "akinori",
		Password: nil,
		Role:     model.Admin,
	}

	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": repo.Put(u),
			"user":    u,
		})
	}
}
