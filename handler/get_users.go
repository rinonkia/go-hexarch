package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/repository"
)

func GetUsers(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		us := userRepository.GetAll()
		successResponse(c, us)
	}
}
