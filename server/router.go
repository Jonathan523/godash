package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (server *Server) setupRouter() {
	server.Router.Get("/", launchpad)
	server.Router.Route("/api", func(r chi.Router) {
		r.Route("/system", func(r chi.Router) {
			r.Get("/static", routeStaticSystem)
			r.Get("/live", routeLiveSystem)
			r.Get("/ws", webSocket)
		})
		r.Get("/weather", getWeather)
	})
	server.serveStatic("static")
	server.serveStatic("storage/icons")
	server.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}
