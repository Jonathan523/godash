package weather

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"godash/hub"
	"io"
	"math"
	"net/http"
	"time"
)

func NewWeatherService(logging *zap.SugaredLogger, hub *hub.Hub) *weather {
	var w = weather{log: logging, hub: hub}
	if err := env.Parse(&w.config); err != nil {
		panic(err)
	}
	if w.config.Key != "" {
		w.setWeatherUnits()
		go w.updateWeather(time.Second * 90)
	}
	return &w
}

func (w *weather) setWeatherUnits() {
	if w.config.Units == "imperial" {
		w.CurrentWeather.Units = "°F"
	} else {
		w.CurrentWeather.Units = "°C"
	}
}

func (w *weather) copyWeatherValues(weatherResp *OpenWeatherApiResponse) {
	myTime := time.Unix(weatherResp.Sys.Sunrise, 0)
	w.CurrentWeather.Sunrise = myTime.Format("15:04")
	myTime = time.Unix(weatherResp.Sys.Sunset, 0)
	w.CurrentWeather.Sunset = myTime.Format("15:04")
	w.CurrentWeather.Icon = weatherResp.Weather[0].Icon
	if w.config.Digits {
		w.CurrentWeather.Temp = weatherResp.Main.Temp
	} else {
		w.CurrentWeather.Temp = math.Round(weatherResp.Main.Temp)
	}
	w.CurrentWeather.Description = weatherResp.Weather[0].Description
	w.CurrentWeather.Humidity = weatherResp.Main.Humidity
}

func (w *weather) updateWeather(interval time.Duration) {
	var weatherResponse OpenWeatherApiResponse
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s&lang=%s",
			w.config.Latitude,
			w.config.Longitude,
			w.config.Key,
			w.config.Units,
			w.config.Lang))
		if err != nil || resp.StatusCode != 200 {
			w.log.Error("weather cannot be updated, please check OPEN_WEATHER_KEY")
		} else {
			body, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &weatherResponse)
			if err != nil {
				w.log.Error("weather cannot be processed")
			} else {
				w.copyWeatherValues(&weatherResponse)
				w.log.Debugw("weather updated", "temp", w.CurrentWeather.Temp)
			}
			resp.Body.Close()
			w.hub.LiveInformationCh <- hub.Message{WsType: hub.Weather, Message: w.CurrentWeather}
		}
		time.Sleep(interval)
	}
}
