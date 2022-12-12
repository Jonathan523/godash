package server

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func (server *Server) setupMiddlewares() {
	server.Router.Use(middleware.RealIP)
	if logrus.GetLevel() == logrus.TraceLevel {
		server.Router.Use(middleware.Logger)
	}
	server.Router.Use(middleware.Recoverer)
	server.Router.Use(middleware.CleanPath)
	server.Router.Use(middleware.RedirectSlashes)
	server.Router.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	server.Router.Use(middleware.Compress(5, "text/html", "text/js", "text/css"))
}
