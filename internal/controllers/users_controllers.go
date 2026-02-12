package controllers

import (
	"net/http"

	"github.com/Quavke/market/internal/services"
)

type UsersController struct {
	service services.UsersService
}

func NewUsersController(service services.UsersService) *UsersController {
	return &UsersController{service: service}
}

func (controller *UsersController) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (controller *UsersController) GetByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}
