package server

import (
	"github.com/go-chi/chi/v5"
	"godash/hub"
)

type Server struct {
	Router  *chi.Mux
	Hub     *hub.Hub
	Port    int
	PageUrl string `mapstructure:"PAGE_URL"`
	Title   string
}
