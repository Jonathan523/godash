package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"godash/bookmarks"
	"godash/system"
	"godash/weather"
	"html/template"
	"net/http"
	"net/url"
)

type goDash struct {
	router *echo.Echo
	logger *zap.SugaredLogger
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

	g.router.Use(middleware.Recover())
	g.router.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	g.router.Pre(middleware.RemoveTrailingSlash())

	w := weather.NewWeatherService(g.logger)
	b := bookmarks.NewBookmarkService(g.logger)
	var s *system.System
	if g.config.LiveSystem {
		s = system.NewSystemService(g.logger)
	}

	g.router.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.gohtml", map[string]interface{}{
			"Title":     g.config.Title,
			"Weather":   w.CurrentWeather,
			"Bookmarks": b.Categories,
			"System":    s,
		})
	})
	g.router.Static("/static", "static")
	g.router.Static("/storage/icons", "storage/icons")

	g.router.GET("/robots.txt", func(c echo.Context) error {
		return c.String(http.StatusOK, "User-agent: *\nDisallow: /")
	})

	g.router.Logger.Fatal(g.router.Start(fmt.Sprintf(":%d", g.config.Port)))
}
