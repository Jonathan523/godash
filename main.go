package main

import (
	"godash/bookmark"
	"godash/logging"
	"godash/server"
	"godash/system"
	"godash/weather"
)

func main() {
	logging.NewGlobalLogger()
	weather.NewWeatherService()
	system.NewSystemService()
	bookmark.NewBookmarkService()
	server.NewServer()
}
