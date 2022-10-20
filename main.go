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
		go w.UpdateWeather(time.Second * 90)
	}
	s := server.NewServer()
	s.Listen()
}
