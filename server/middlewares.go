package server

import (
	"github.com/hertz-contrib/gzip"
)

func (server *Server) setupMiddlewares() {
	server.Router.Use(gzip.Gzip(gzip.DefaultCompression))
}
