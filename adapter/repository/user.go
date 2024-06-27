package repository

import "github.com/rinonkia/go-hexagonal-architecture/core/model"

type inMemoryUserRepository struct {
	users map[string]*model.User
}

func NewInMemoryUserRepository() *inMemoryUserRepository {
	return &inMemoryUserRepository{users: map[string]*model.User{}}
}

func (r *inMemoryUserRepository) Put(u *model.User) bool {
	r.users[u.Name] = u
	return true
}

func (r *inMemoryUserRepository) GetAll() []*model.User {
	us := make([]*model.User, 0, len(r.users))
	for _, v := range r.users {
		us = append(us, v)
	}
	return us
}
