package middleware

import (
	"github.com/rinonkia/go-hexarch/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const tokenHeaderKey = "Authorization"

func CheckAuthorization(tokenService *service.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := tokenService.CheckToken(c.GetHeader(tokenHeaderKey)); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
}
