package main

func (g *goDash) setupRouter() {
	g.router.GET("/", g.index)
	g.router.GET("/ws", g.ws)
	g.router.GET("/robots.txt", robots)

	static := g.router.Group("/static")
	static.Use(longCacheLifetime)
	static.Static("/", "static")

	storage := g.router.Group("/storage")
	storage.Use(longCacheLifetime)
	storage.Static("/icons", "storage/icons")

	g.router.RouteNotFound("/*", redirectHome)
}
