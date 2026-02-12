package users

import "time"

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateReq struct {
	Username string `json:"username"`
}

type UserDB struct {
	ID           int
	Username     string
	PasswordHash []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TODO продавец или нет

type UserResponse struct {
	Username string `json:"username"`
}
