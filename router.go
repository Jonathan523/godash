package main

func (g *goDash) setupRouter() {
	g.router.GET("/", g.index)
	g.router.GET("/ws", g.ws)
	g.router.GET("/robots.txt", robots)
	g.router.Static("/static", "static")
	storage := g.router.Group("/storage")
	storage.Use(longCacheLifetime)
	storage.Static("/icons", "storage/icons")
	g.router.RouteNotFound("/*", redirectHome)
}
