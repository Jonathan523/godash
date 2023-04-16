package main

import (
	"godash/bookmarks"
	"godash/hub"
	"godash/system"
	"godash/weather"

	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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
	bookmarks *bookmarks.Config
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

func (g *goDash) startServer() {
	if err := g.router.Start(fmt.Sprintf(":%d", g.config.Port)); err != nil && err != http.ErrServerClosed {
		g.logger.Fatal("shutting down the server")
	}
}

func (g *goDash) setupTemplateRender() {
	g.router.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func main() {
	g := goDash{router: echo.New()}
	if err := env.Parse(&g.config); err != nil {
		panic(err)
	}

	g.setupTemplateRender()
	g.setupLogger()
	defer func(logger *zap.SugaredLogger) {
		_ = logger.Sync()
	}(g.logger)
	g.setupEchoLogging()
	g.setupMiddlewares()
	g.createInfoServices()
	g.setupRouter()

	go g.startServer()
	g.logger.Infof("running on %s:%d", "http://localhost", g.config.Port)

	quit := make(chan os.Signal, 1)
	// https://docs.docker.com/engine/reference/commandline/stop/
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := g.router.Shutdown(ctx); err != nil {
		g.logger.Fatal(err)
	}
}
