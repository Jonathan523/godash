package server

import (
	"github.com/go-chi/chi/v5"
	"godash/files"
	"net/http"
)

func (server *Server) setupRouter() {
	server.Router.Get("/", launchpad)
	server.Router.Route("/api", func(r chi.Router) {
		r.Get("/ws", webSocket)
		r.Get("/weather", getWeather)
	})
	server.Router.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, files.TemplatesFolder+"/robots.txt")
	})
	server.Router.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon/favicon.ico")
	})
	server.serveStatic("static")
	server.serveStatic("storage/icons")
	server.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}
