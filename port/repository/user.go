package repository

import (
	"github.com/rinonkia/go-hexarch/core/model"
	"github.com/rs/xid"
)

type UserRepository interface {
	Put(user *model.User) bool
	GetByID(id xid.ID) (*model.User, error)
	GetAll() []*model.User
}
