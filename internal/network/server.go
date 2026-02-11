package network

import (
	"net/http"
	"os"

	"github.com/Quavke/market/internal/controllers"
	"github.com/Quavke/market/internal/repositories"
	"github.com/Quavke/market/internal/routes"
	"github.com/Quavke/market/internal/services"
)

type server struct {
	cfg  *serverConfig
	mux  *http.ServeMux
	stop chan struct{}
}

type serverConfig struct {
	Port   string
	DBHost string
}

func NewServerConfig() *serverConfig {
	portStr := ":" + os.Getenv("Port")
	if portStr == ":" {
		portStr = ":8080"
	}
	return &serverConfig{
		Port: portStr,
	}

}

func NewServer(cfg *serverConfig) *server {
	all_handlers := make(map[string]*http.HandlerFunc, 0)

	// TODO: add users stuff
	usersRepo := repositories.NewUsersRepository(cfg.DBHost)
	usersService := services.NewUsersService(usersRepo)
	usersController := controllers.NewUsersController(usersService)
	usersRouter := routes.NewUsersRouter(usersController).(map[string]*http.HandlerFunc)
	all_handlers = append(all_handlers, usersRouter)

	///...

	mux := http.NewServeMux()

	if len(all_handlers) > 0 {
		for route, handler := range all_handlers {
			mux.HandleFunc(route, *handler)
		}
	} else {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("There is nothing")) })
	}

	return &server{
		cfg:  cfg,
		mux:  mux,
		stop: make(chan struct{}),
	}
}

func (s *server) Start() error {
	return http.ListenAndServe(s.cfg.Port, s.mux)
}

func (s *server) Stop() {
	close(s.stop)
}
