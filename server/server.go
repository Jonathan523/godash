package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"launchpad/message"
	"net/http"
)

type Server struct {
	Router *chi.Mux
	Port   int
}

func NewServer(port int) *Server {
	router := chi.NewRouter()
	setupMiddlewares(router)
	setupRouter(router)
	return &Server{router, port}
}

func (server *Server) Listen() {
	logrus.WithField("port", server.Port).Info("application running")
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), server.Router)
	if err != nil {
		logrus.WithField("error", err).Fatal(message.CannotStart.String())
	}
}
