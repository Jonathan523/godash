package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func setupMiddlewares(router *chi.Mux) {
	if logrus.GetLevel() == logrus.TraceLevel {
		logger := logrus.New()
		logger.Formatter = &logrus.TextFormatter{TimestampFormat: "2006/01/02 15:04:05", FullTimestamp: true}
		router.Use(newStructuredLogger(logger))
	}
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	router.Use(middleware.RealIP)
	router.Use(middleware.CleanPath)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Compress(5, "text/html", "text/css"))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept-Encoding", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}
