package server

import (
	"github.com/go-chi/chi/v5"
)

func setupRouter(router *chi.Mux) {
	router.Get("/", launchpad)
	router.Get("/static/style.css", serveStyleCss)
}
