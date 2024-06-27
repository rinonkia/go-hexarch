package model

import "github.com/rs/xid"

type (
	User struct {
		ID       xid.ID
		Name     string
		Password []byte
		Role     Role
	}

	Role string
)

const (
	General Role = "general"
	Admin   Role = "admin"
)
