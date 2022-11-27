package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/gzip"
)

func CacheMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Header("Cache-Control", "public, max-age=604800, immutable")
	}
}
func (server *Server) setupMiddlewares() {
	server.Router.Use(gzip.Gzip(gzip.DefaultCompression))
}
