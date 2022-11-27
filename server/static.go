package server

import (
	"github.com/cloudwego/hertz/pkg/app"
)

func (server *Server) serveStatic(folder string) {
	server.Router.StaticFS("/"+folder, &app.FS{Root: "./", GenerateIndexPages: true})
}
