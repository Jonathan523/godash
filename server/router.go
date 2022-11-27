package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func (server *Server) setupRouter() {
	server.Router.GET("/", server.goDash)
	server.Router.GET("/ws", webSocket)

	server.Router.NoMethod(func(c context.Context, ctx *app.RequestContext) {
		ctx.Redirect(consts.StatusPermanentRedirect, []byte("/"))
	})
	server.Router.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.Redirect(consts.StatusPermanentRedirect, []byte("/"))
	})

	server.Router.Use(CacheMiddleware())
	server.serveStatic("static")
	server.serveStatic("storage/icons")

	server.Router.GET("/robots.txt", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "User-agent: *\nDisallow: /")
	})
}
