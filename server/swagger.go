package server

import (
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"godash/docs"
	"net/http"
	"net/url"
)

func redirectToSwagger(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger/index.html", http.StatusTemporaryRedirect)
}

func (server *Server) setupSwagger() {
	if server.Swagger {
		docs.SwaggerInfo.Title = "GoDash"
		docs.SwaggerInfo.Version = "1.0.0"
		docs.SwaggerInfo.BasePath = "/api"
		parsed, _ := url.Parse(server.AllowedHosts[0])
		docs.SwaggerInfo.Host = parsed.Host

		server.Router.Get("/swagger", redirectToSwagger)
		server.Router.Get("/swagger/", redirectToSwagger)
		server.Router.Get("/swagger/*", httpSwagger.Handler())
		logrus.WithField("url", server.AllowedHosts[0]+"/swagger").Info("swagger running")
	}
}
