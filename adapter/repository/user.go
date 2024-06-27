package repository

import (
	"errors"

	"github.com/rinonkia/go-hexarch/core/model"
	"github.com/rs/xid"
)

type users map[xid.ID]*model.User

type inMemoryUserRepository struct {
	im users
}

func NewInMemoryUserRepository() *inMemoryUserRepository {
	return &inMemoryUserRepository{
		im: users{},
	}
}

func (r *inMemoryUserRepository) Put(u *model.User) bool {
	r.im[u.ID] = u
	return true
}

func (r *inMemoryUserRepository) GetByID(id xid.ID) (*model.User, error) {
	u := r.im[id]
	if u == nil {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *inMemoryUserRepository) GetAll() []*model.User {
	us := make([]*model.User, 0, len(r.im))
	for _, v := range r.im {
		us = append(us, v)
	}
	return us
}
