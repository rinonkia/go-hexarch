package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/domain/entity"
	"github.com/rinonkia/go-hexarch/domain/service"
	"github.com/rinonkia/go-hexarch/repository"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(
	tokenService *service.Token,
	userRepository repository.UserRepository,
) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		un := c.PostForm("username")
		pw := c.PostForm("password")
		_ = c.PostForm("role")

		p, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		u := &entity.User{
			ID:       xid.New(),
			Name:     un,
			Password: p,
			Role:     entity.Admin,
		}

		token, err := tokenService.GenerateToken(u.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": userRepository.Put(u),
			"user":    u,
			"token":   token,
		})
	}
}
