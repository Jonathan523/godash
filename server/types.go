package server

import (
	"github.com/go-chi/chi/v5"
	"godash/hub"
)

type Server struct {
	Router       *chi.Mux
	Hub          *hub.Hub
	Port         int
	AllowedHosts []string `mapstructure:"ALLOWED_HOSTS"`
	Title        string
}
