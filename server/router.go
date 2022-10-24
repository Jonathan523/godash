package server

import (
	"godash/files"
	"net/http"
)

func (server *Server) setupRouter() {
	server.Router.Get("/", launchpad)
	server.Router.Get("/ws", webSocket)
	
	server.serveStatic("static")
	server.serveStatic("storage/icons")

	server.Router.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, files.TemplatesFolder+"/robots.txt")
	})
	server.Router.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon/favicon.ico")
	})

	server.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}
