package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func (server *Server) setupRouter() {
	server.Router.GET("/", server.goDash)
	server.Router.GET("/ws", webSocket)

	server.serveStatic("static")
	server.serveStatic("storage/icons")

	server.Router.GET("/robots.txt", func(c context.Context, ctx *app.RequestContext) {
		ctx.File(TemplatesFolder + "/robots.txt")
	})
	server.Router.GET("/favicon.ico", func(c context.Context, ctx *app.RequestContext) {
		ctx.File("static/favicon/favicon.ico")
	})

	server.Router.NoMethod(func(c context.Context, ctx *app.RequestContext) {
		ctx.Redirect(consts.StatusPermanentRedirect, []byte("/"))
	})
	server.Router.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.Redirect(consts.StatusPermanentRedirect, []byte("/"))
	})
}
