package server

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func (server *Server) setupMiddlewares() {
	server.Router.Use(middleware.RealIP)
	if logrus.GetLevel() == logrus.TraceLevel {
		logger := logrus.New()
		logger.Formatter = &logrus.TextFormatter{TimestampFormat: "2006/01/02 15:04:05", FullTimestamp: true}
		server.Router.Use(newStructuredLogger(logger))
	}
	server.Router.Use(middleware.Recoverer)
	server.Router.Use(middleware.CleanPath)
	server.Router.Use(middleware.RedirectSlashes)
	server.Router.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	server.Router.Use(middleware.Compress(5, "text/html", "text/js", "text/css"))
	server.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{server.PageUrl},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept-Encoding", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}
