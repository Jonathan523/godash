package main

import (
	"godash/logging"
	"godash/server"
	"godash/weather"
)

func main() {
	logging.NewGlobalLogger()
	weather.NewWeather()
	server.NewServer()
}
