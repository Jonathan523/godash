package main

import (
	"godash/logging"
	"godash/server"
	"godash/weather"
	"time"
)

func main() {
	logging.NewGlobalLogger()
	w := weather.NewWeather()
	if w.OpenWeather.Key != "" {
		w.SetWeatherUnits()
		go w.UpdateWeather(time.Second * 150)
	}
	s := server.NewServer()
	s.Listen()
}
