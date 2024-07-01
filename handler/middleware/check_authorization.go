package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rinonkia/go-hexarch/config"
)

var InvalidTokenError = errors.New("invalid token")

const tokenHeaderKey = "Authorization"

func CheckAuthorization(sk config.SecretKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		var claim jwt.RegisteredClaims
		tokenString, err := extractToken(c.GetHeader(tokenHeaderKey))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(sk), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "token invalid",
			})
			return
		}
	}
}

func extractToken(s string) (string, error) {
	if s == "" {
		return "", InvalidTokenError
	}

	if len(strings.Split(s, " ")) != 2 {
		return "", InvalidTokenError
	}

	split := strings.Split(s, " ")
	if split[0] != "Bearer" {
		return "", InvalidTokenError
	}

	return split[1], nil
}
