package service

import "github.com/rs/xid"

type (
	TokenGenerator interface {
		Exec(dto *TokenGeneratorDTO) *TokenGeneratorResult
	}

	TokenGeneratorDTO struct {
		ID xid.ID
	}

	TokenGeneratorResult struct {
		Err   error
		Token string
	}
)
