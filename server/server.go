package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"godash/config"
	"godash/message"
	"net/http"
)

type Server struct {
	Router       *chi.Mux
	Port         int
	AllowedHosts []string `mapstructure:"ALLOWED_HOSTS"`
}

func NewServer() *Server {
	server := Server{}
	config.ParseViperConfig(&server, config.AddViperConfig("server"))
	server.Router = chi.NewRouter()
	setupMiddlewares(server.Router)
	setupRouter(server.Router)
	return &server
}

func (server *Server) Listen() {
	logrus.WithField("port", server.Port).Info("application running")
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), server.Router)
	if err != nil {
		logrus.WithField("error", err).Fatal(message.CannotStart.String())
	}
}
