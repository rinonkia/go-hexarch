package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/core/model"
	"github.com/rinonkia/go-hexarch/port/repository"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(repo repository.UserRepository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		un := ctx.PostForm("username")
		pw := ctx.PostForm("password")
		_ = ctx.PostForm("role")

		p, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		u := &model.User{
			ID:       xid.New(),
			Name:     un,
			Password: p,
			Role:     model.Admin,
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": repo.Put(u),
			"user":    u,
		})
	}
}
