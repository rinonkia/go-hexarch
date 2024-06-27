package model

type (
	User struct {
		Name     string
		Password []byte
		Role
	}

	Role string
)

const (
	General Role = "general"
	Admin   Role = "admin"
)
