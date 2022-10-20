package weather

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"godash/config"
	"io"
	"net/http"
	"time"
)

var Conf = Config{}
var CurrentOpenWeather = OpenWeatherApiResponse{}

func NewWeather() {
	config.ParseViperConfig(&Conf, config.AddViperConfig("weather"))
	if Conf.OpenWeather.Key != "" {
		setWeatherUnits()
		go updateWeather(time.Second * 150)
	}
}

func setWeatherUnits() {
	if Conf.OpenWeather.Units == "imperial" {
		CurrentOpenWeather.Units = "°F"
	} else {
		CurrentOpenWeather.Units = "°C"
	}
}

func calcTimestamps() {
	myTime := time.Unix(CurrentOpenWeather.Sys.Sunrise, 0)
	CurrentOpenWeather.Sys.StrSunrise = myTime.Format("15:04")
	myTime = time.Unix(CurrentOpenWeather.Sys.Sunset, 0)
	CurrentOpenWeather.Sys.StrSunset = myTime.Format("15:04")
}

func updateWeather(interval time.Duration) {
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s&lang=%s",
			Conf.Location.Latitude,
			Conf.Location.Longitude,
			Conf.OpenWeather.Key,
			Conf.OpenWeather.Units,
			Conf.OpenWeather.Lang))
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
			calcTimestamps()
			resp.Body.Close()
		}
		time.Sleep(interval)
	}
}
