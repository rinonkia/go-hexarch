package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/port/repository"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func Login(repo repository.UserRepository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.PostForm("id")

		x, err := xid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		u, err := repo.GetByID(x)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		pw := ctx.PostForm("password")
		if err = bcrypt.CompareHashAndPassword(u.Password, []byte(pw)); err != nil {
			status := http.StatusInternalServerError
			if errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) {
				status = http.StatusNotFound
			}

			ctx.JSON(status, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	}
}
