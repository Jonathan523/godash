package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func longCacheLifetime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderCacheControl, "public, max-age=31536000, immutable")
		return next(c)
	}
}

func (g *goDash) setupMiddlewares() {
	g.router.Use(middleware.Recover())
	g.router.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	g.router.Pre(middleware.RemoveTrailingSlash())
	g.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: g.config.AllowedHosts,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderCacheControl},
		AllowMethods: []string{echo.GET, http.MethodHead},
	}))
}
