package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/repository"
)

func GetUsers(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		us := userRepository.GetAll()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"users":   us,
		})
	}
}
