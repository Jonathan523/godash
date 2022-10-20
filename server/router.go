package server

import (
	"github.com/go-chi/chi/v5"
)

func (server *Server) setupRouter() {
	server.Router.Get("/", launchpad)
	server.Router.Route("/system", func(r chi.Router) {
		r.Get("/static", routeStaticSystem)
		r.Get("/live", routeLiveSystem)
		r.Get("/ws", webSocket)
	})
	server.serveStatic("static")
	server.serveStatic("storage/icons")
}
