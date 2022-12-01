package server

import (
	"fmt"
	hertz "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/sirupsen/logrus"
	"godash/config"
	"godash/hub"
)

var server = Server{}

const TemplatesFolder = "templates/"

func NewServer() {
	config.ParseViperConfig(&server, config.AddViperConfig("server"))
	server.Router = hertz.Default(
		hertz.WithHostPorts(fmt.Sprintf(":%d", server.Port)),
		hertz.WithRemoveExtraSlash(true),
		hertz.WithRedirectTrailingSlash(true),
		hertz.WithGetOnly(true),
		hertz.WithAutoReloadRender(true, 0),
	)
	setupLogging()
	server.setupMiddlewares()
	server.prepareHtml()
	server.Hub = hub.NewHub()
	server.setupRouter()
	server.Listen()
}

func (server *Server) prepareHtml() {
	server.Router.LoadHTMLGlob(TemplatesFolder + "*")
}

func (server *Server) Listen() {
	logrus.WithField("port", server.Port).Info("server starting")
	server.Router.Spin()
}
