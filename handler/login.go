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

var userDoesNotMatchErr = errors.New("user or/and password does not match")

func Login(
	tokenService *service.Token,
	userRepository repository.UserRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.PostForm("id")

		x, err := xid.FromString(id)
		if err != nil {
			failedResponse(c, http.StatusBadRequest, err)
			return
		}
		u, err := userRepository.GetByID(x)
		if err != nil {
			failedResponse(c, http.StatusBadRequest, userDoesNotMatchErr)
			return
		}

		pw := c.PostForm("password")
		if err = bcrypt.CompareHashAndPassword(u.Password, []byte(pw)); err != nil {
			if errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) {
				failedResponse(c, http.StatusBadRequest, userDoesNotMatchErr)
				return
			}

			failedResponse(c, http.StatusInternalServerError, err)
			return
		}

		token, err := tokenService.GenerateToken(u.ID)
		if err != nil {
			failedResponse(c, http.StatusInternalServerError, err)
			return
		}

		successResponse(c, token)
	}
}
