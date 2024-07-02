package im

import (
	"errors"

	"github.com/rinonkia/go-hexarch/domain/entity"
	"github.com/rs/xid"
)

type users map[xid.ID]*entity.User

type InMemoryUserRepository struct {
	im users
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		im: users{},
	}
}

func (r *InMemoryUserRepository) Put(u *entity.User) bool {
	r.im[u.ID] = u
	return true
}

func (r *InMemoryUserRepository) GetByID(id xid.ID) (*entity.User, error) {
	u := r.im[id]
	if u == nil {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *InMemoryUserRepository) GetAll() []*entity.User {
	us := make([]*entity.User, 0, len(r.im))
	for _, v := range r.im {
		us = append(us, v)
	}
	return us
}
