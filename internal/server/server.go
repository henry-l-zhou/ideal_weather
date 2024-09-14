package server

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	goweather "github.com/henry-l-zhou/ideal_weather/go-weather"
)

type Server struct {
	router        *chi.Mux
	weatherClient goweather.Client
}

func NewServer() *Server {
	weatherApiKey := os.Getenv("OPENWEATHER_API_KEY")

	s := &Server{
		weatherClient: *goweather.NewClient(weatherApiKey),
		router:        chi.NewRouter(),
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
