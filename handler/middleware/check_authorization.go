package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/domain/service"
)

const tokenHeaderKey = "Authorization"

func CheckAuthorization(tokenService *service.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := tokenService.CheckToken(c.GetHeader(tokenHeaderKey)); err != nil {
			failedAuthResponse(c, err)
			return
		}
	}
}

func failedAuthResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"data":  nil,
		"error": err.Error(),
	})

}
