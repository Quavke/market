package services

import (
	"github.com/Quavke/market/internal/models/users"
	"github.com/Quavke/market/internal/repositories"
)

type UsersService interface {
	GetAll() (*users.UserResponse, error)
	GetByID(id int) (*users.UserResponse, error)
}

type UsersServiceImpl struct {
	repo repositories.UsersRepository
}

func NewUsersServiceImpl(repo repositories.UsersRepository) UsersService {
	return &UsersServiceImpl{repo: repo}
}

func (u *UsersServiceImpl) GetAll() (*users.UserResponse, error) {
	return nil, nil
}

func (u *UsersServiceImpl) GetByID(id int) (*users.UserResponse, error) {
	return nil, nil
}
