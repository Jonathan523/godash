package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"godash/bookmarks"
	"godash/hub"
	"godash/system"
	"godash/weather"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type goDash struct {
	router *echo.Echo
	logger *zap.SugaredLogger
	hub    *hub.Hub
	config config
	info   info
}

type info struct {
	weather   *weather.Weather
	bookmarks *bookmarks.Bookmarks
	system    *system.System
}

type config struct {
	Title        string   `env:"TITLE" envDefault:"goDash"`
	Port         int      `env:"PORT" envDefault:"4000"`
	AllowedHosts []string `env:"ALLOWED_HOSTS" envDefault:"*" envSeparator:","`
	LogLevel     string   `env:"LOG_LEVEL" envDefault:"info"`
	LiveSystem   bool     `env:"LIVE_SYSTEM" envDefault:"true"`
}

func (g *goDash) createInfoServices() {
	g.hub = hub.NewHub(g.logger)
	g.info = info{
		weather:   weather.NewWeatherService(g.logger, g.hub),
		bookmarks: bookmarks.NewBookmarkService(g.logger),
		system:    system.NewSystemService(g.config.LiveSystem, g.logger, g.hub),
	}
}

func (g *goDash) setupMiddlewares() {
	g.router.Use(middleware.Recover())
	g.router.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	g.router.Pre(middleware.RemoveTrailingSlash())
	g.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: g.config.AllowedHosts,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET},
	}))
}

func main() {
	g := goDash{router: echo.New()}
	g.router.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}
	if err := env.Parse(&g.config); err != nil {
		panic(err)
	}

	g.setupLogger()
	defer func(logger *zap.SugaredLogger) {
		_ = logger.Sync()
	}(g.logger)
	g.setupEchoLogging()
	g.setupMiddlewares()
	g.createInfoServices()
	g.router.GET("/", g.index)
	g.router.GET("/ws", g.ws)
	g.router.GET("/robots.txt", robots)
	g.router.Static("/static", "static")
	g.router.Static("/storage/icons", "storage/icons")
	g.router.RouteNotFound("/*", redirectHome)

	go func() {
		if err := g.router.Start(fmt.Sprintf(":%d", g.config.Port)); err != nil && err != http.ErrServerClosed {
			g.logger.Fatal("shutting down the server")
		}
	}()
	g.logger.Infof("running on %s:%d", "http://localhost", g.config.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := g.router.Shutdown(ctx); err != nil {
		g.logger.Fatal(err)
	}
}
