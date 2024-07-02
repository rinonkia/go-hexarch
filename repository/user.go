package repository

import (
	"github.com/rinonkia/go-hexarch/domain/entity"
	"github.com/rs/xid"
)

type UserRepository interface {
	Put(user *entity.User) bool
	GetByID(id xid.ID) (*entity.User, error)
	GetAll() []*entity.User
}
