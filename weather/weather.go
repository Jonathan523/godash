package weather

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"godash/config"
	"godash/hub"
	"io"
	"net/http"
	"time"
)

var CurrentOpenWeather = OpenWeatherApiResponse{}

func NewWeather() *Config {
	conf := Config{}
	config.ParseViperConfig(&conf, config.AddViperConfig("weather"))
	return &conf
}

func (conf *Config) SetWeatherUnits() {
	if conf.OpenWeather.Units == "imperial" {
		CurrentOpenWeather.Units = "°F"
	} else {
		CurrentOpenWeather.Units = "°C"
	}
}

func calcWeatherTimestamps() {
	myTime := time.Unix(CurrentOpenWeather.Sys.Sunrise, 0)
	CurrentOpenWeather.Sys.StrSunrise = myTime.Format("15:04")
	myTime = time.Unix(CurrentOpenWeather.Sys.Sunset, 0)
	CurrentOpenWeather.Sys.StrSunset = myTime.Format("15:04")
}

func (conf *Config) UpdateWeather(interval time.Duration) {
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s&lang=%s",
			conf.Location.Latitude,
			conf.Location.Longitude,
			conf.OpenWeather.Key,
			conf.OpenWeather.Units,
			conf.OpenWeather.Lang))
		if err != nil {
			logrus.Error("weather cannot be updated")
		} else {
			body, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &CurrentOpenWeather)
			if err != nil {
				logrus.Error("weather cannot be processed")
			} else {
				logrus.WithFields(logrus.Fields{"temp": fmt.Sprintf("%0.2f%s", CurrentOpenWeather.Main.Temp, CurrentOpenWeather.Units), "location": CurrentOpenWeather.Name}).Trace("weather updated")
			}
			calcWeatherTimestamps()
			resp.Body.Close()
		}
		hub.LiveInformationCh <- hub.Message{WsType: hub.Weather, Message: CurrentOpenWeather}
		time.Sleep(interval)
	}
}
