package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rinonkia/go-hexarch/config"
	"github.com/rinonkia/go-hexarch/interface/service"
)

type tokenGeneratorImpl struct {
	secretKey []byte
}

func NewTokenGenerator(sk config.SecretKey) service.TokenGenerator {
	return &tokenGeneratorImpl{
		secretKey: []byte(sk),
	}
}

func (s *tokenGeneratorImpl) Exec(dto *service.TokenGeneratorDTO) *service.TokenGeneratorResult {
	var result service.TokenGeneratorResult
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    "https://example.com",
		Subject:   "Access Token",
		ExpiresAt: &jwt.NumericDate{Time: now.Add(time.Minute * 10)},
		IssuedAt:  &jwt.NumericDate{Time: now},
		ID:        dto.ID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(s.secretKey)
	if err != nil {
		result.Err = err
		return &result
	}

	result.Token = t
	return &result
}
