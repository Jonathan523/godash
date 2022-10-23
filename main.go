package main

import (
	"godash/logging"
	"godash/server"
	"godash/system"
	"godash/weather"
)

func main() {
	logging.NewGlobalLogger()
	weather.NewWeather()
	system.NewSystem()
	server.NewServer()
}
