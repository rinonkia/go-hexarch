package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/adapter/service"
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
