package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"godash/bookmarks"
	"godash/weather"
	"html/template"
	"net/http"
	"net/url"
)

type goDash struct {
	router *echo.Echo
	logger *zap.Logger
	config config
}

type config struct {
	Title      string  `env:"TITLE" envDefault:"goDash"`
	Port       int     `env:"PORT" envDefault:"4000"`
	PageUrl    url.URL `env:"PAGE_URL" envDefault:"http://localhost:4000"`
	LogLevel   string  `env:"LOG_LEVEL" envDefault:"info"`
	LiveSystem bool    `env:"LIVE_SYSTEM" envDefault:"true"`
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
	defer g.logger.Sync()
	g.setupEchoLogging()

	w := weather.NewWeatherService(g.logger.Sugar())
	b := bookmarks.NewBookmarkService(g.logger.Sugar())

	g.router.Use(middleware.Recover())
	g.router.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	g.router.Pre(middleware.RemoveTrailingSlash())

	g.router.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.gohtml", map[string]interface{}{
			"Title":     g.config.Title,
			"Weather":   w.CurrentWeather,
			"Bookmarks": b.Categories,
		})
	})
	g.router.Static("/static", "static")
	g.router.Logger.Fatal(g.router.Start(fmt.Sprintf(":%d", g.config.Port)))
}
