package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/domain/service"
	"github.com/rinonkia/go-hexarch/repository"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func Login(
	tokenService *service.Token,
	userRepository repository.UserRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.PostForm("id")

		x, err := xid.FromString(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		u, err := userRepository.GetByID(x)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		pw := c.PostForm("password")
		if err = bcrypt.CompareHashAndPassword(u.Password, []byte(pw)); err != nil {
			status := http.StatusInternalServerError
			if errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) {
				status = http.StatusNotFound
			}

			c.JSON(status, gin.H{
				"message": err.Error(),
			})
			return
		}

		token, err := tokenService.GenerateToken(u.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":  u,
			"token": token,
		})
	}
}
