package network

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Quavke/market/internal/controllers"
	"github.com/Quavke/market/internal/repositories"
	"github.com/Quavke/market/internal/routes"
	"github.com/Quavke/market/internal/services"
	"github.com/Quavke/market/pkg/db"
)

type Server struct {
	cfg  *ServerConfig
	mux  *http.ServeMux
	stop chan struct{}
}

type ServerConfig struct {
	Port   string
	DBHost string
}

func NewServerConfig() *ServerConfig {
	portStr := ":" + os.Getenv("Port")
	if portStr == ":" {
		portStr = ":8080"
	}
	return &ServerConfig{
		Port: portStr,
	}

}

func NewServer(cfg *ServerConfig) (*Server, error) {
	allHandlers := make(map[string]http.HandlerFunc)

	pgDB, err := db.Connect()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return nil, err
	}

	// TODO: add users stuff
	usersRepo := repositories.NewUsersPGRepository(pgDB)
	usersService := services.NewUsersServiceImpl(usersRepo)
	usersController := controllers.NewUsersController(usersService)
	usersRouter := routes.NewUsersRouter(usersController)
	for path, handler := range usersRouter {
		allHandlers[path] = handler
	}

	///...

	mux := http.NewServeMux()

	if len(allHandlers) == 0 {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("There is nothing")) })
	}

	return &Server{
		cfg:  cfg,
		mux:  mux,
		stop: make(chan struct{}),
	}, nil
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.cfg.Port, s.mux)
}

func (s *Server) Stop() {
	close(s.stop)
}
