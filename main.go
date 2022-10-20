package main

import (
	"launchpad/logging"
	"launchpad/server"
	"launchpad/weather"
)

func main() {
	logging.NewGlobalLogger()
	weather.NewWeather()
	s := server.NewServer()
	s.Listen()
}
