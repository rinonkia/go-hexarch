package repository

import "github.com/rinonkia/go-hexagonal-architecture/core/model"

type UserRepository interface {
	Put(user *model.User) bool
	GetAll() []*model.User
}
