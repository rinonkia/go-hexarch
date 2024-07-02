package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rinonkia/go-hexarch/config"
	"github.com/rs/xid"
)

type Token struct {
	sk []byte
}

func NewToken(sk config.SecretKey) *Token {
	return &Token{
		sk: []byte(sk),
	}
}

var invalidTokenErr = errors.New("invalid token")

func (s *Token) GenerateToken(id xid.ID) (token string, err error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    "https://example.com",
		Subject:   "Access Token",
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 10)),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        id.String(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(s.sk)
	return
}

func (s *Token) CheckToken(t string) error {
	tokenString, err := extractToken(t)
	if err != nil {
		return err
	}

	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.sk, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return invalidTokenErr
	}

	return nil
}

func extractToken(s string) (string, error) {
	if s == "" {
		return "", invalidTokenErr
	}

	if len(strings.Split(s, " ")) != 2 {
		return "", invalidTokenErr
	}

	split := strings.Split(s, " ")
	if split[0] != "Bearer" {
		return "", invalidTokenErr
	}

	return split[1], nil
}
