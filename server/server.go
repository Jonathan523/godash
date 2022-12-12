package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"godash/config"
	"godash/hub"
	"godash/message"
	"net/http"
)

var server = Server{}

func NewServer() {
	config.ParseViperConfig(&server, config.AddViperConfig("server"))
	server.Router = chi.NewRouter()
	server.Hub = hub.NewHub()
	server.setupMiddlewares()
	server.setupRouter()
	server.Listen()
}

func (server *Server) Listen() {
	logrus.WithField("port", server.Port).Info("application running")
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), server.Router)
	if err != nil {
		logrus.WithField("error", err).Fatal(message.CannotStart.String())
	}
}
