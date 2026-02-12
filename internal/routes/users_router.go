package routes

import (
	"github.com/Quavke/market/internal/controllers"
	"net/http"
)

func NewUsersRouter(ctrl *controllers.UsersController) map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/users/":    ctrl.GetAll,
		"/users/:id": ctrl.GetByID,
	}
}
