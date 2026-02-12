package repositories

import (
	"database/sql"

	"github.com/Quavke/market/internal/models/users"
)

type UsersRepository interface {
	GetAll() (*users.UserResponse, error)
	GetByID(id int) (*users.UserResponse, error)
}

type UsersPGRepository struct {
	db *sql.DB
}

func NewUsersPGRepository(db *sql.DB) UsersRepository {
	return &UsersPGRepository{db: db}
}

func (u *UsersPGRepository) GetAll() (*users.UserResponse, error) {
	return nil, nil
}

func (u *UsersPGRepository) GetByID(id int) (*users.UserResponse, error) {
	return nil, nil
}
