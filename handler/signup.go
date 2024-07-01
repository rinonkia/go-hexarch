package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/core/model"
	"github.com/rinonkia/go-hexarch/interface/repository"
	"github.com/rinonkia/go-hexarch/interface/service"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(
	tokenGenerator service.TokenGenerator,
	userRepository repository.UserRepository,
) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		un := c.PostForm("username")
		pw := c.PostForm("password")
		_ = c.PostForm("role")

		p, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
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

		sr := tokenGenerator.Exec(&service.TokenGeneratorDTO{
			ID: u.ID,
		})
		if sr.Err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": sr.Err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": userRepository.Put(u),
			"user":    u,
			"token":   sr.Token,
		})
	}
}
