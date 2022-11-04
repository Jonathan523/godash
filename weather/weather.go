package weather

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"godash/config"
	"godash/hub"
	"io"
	"math"
	"net/http"
	"time"
)

var Conf = PackageConfig{}
var CurrentWeather = Weather{}

func NewWeatherService() {
	config.ParseViperConfig(&Conf, config.AddViperConfig("weather"))
	if Conf.OpenWeather.Key != "" {
		setWeatherUnits()
		go updateWeather(time.Second * 90)
	}
}

func setWeatherUnits() {
	if Conf.OpenWeather.Units == "imperial" {
		CurrentWeather.Units = "°F"
	} else {
		CurrentWeather.Units = "°C"
	}
}

func copyWeatherValues(weatherResp *OpenWeatherApiResponse) {
	myTime := time.Unix(weatherResp.Sys.Sunrise, 0)
	CurrentWeather.Sunrise = myTime.Format("15:04")
	myTime = time.Unix(weatherResp.Sys.Sunset, 0)
	CurrentWeather.Sunset = myTime.Format("15:04")
	CurrentWeather.Icon = weatherResp.Weather[0].Icon
	if Conf.OpenWeather.Digits {
		CurrentWeather.Temp = weatherResp.Main.Temp
	} else {
		CurrentWeather.Temp = math.Round(weatherResp.Main.Temp)
	}
	CurrentWeather.Description = weatherResp.Weather[0].Description
	CurrentWeather.Humidity = weatherResp.Main.Humidity
}

func updateWeather(interval time.Duration) {
	var weatherResponse OpenWeatherApiResponse
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s&lang=%s",
			Conf.Location.Latitude,
			Conf.Location.Longitude,
			Conf.OpenWeather.Key,
			Conf.OpenWeather.Units,
			Conf.OpenWeather.Lang))
		if err != nil || resp.StatusCode != 200 {
			logrus.Error("weather cannot be updated, please check OPEN_WEATHER_KEY")
		} else {
			body, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &weatherResponse)
			if err != nil {
				logrus.Error("weather cannot be processed")
			} else {
				copyWeatherValues(&weatherResponse)
				logrus.WithFields(logrus.Fields{"temp": fmt.Sprintf("%0.2f%s", CurrentWeather.Temp, CurrentWeather.Units), "humidity": fmt.Sprintf("%d%s", CurrentWeather.Humidity, "%")}).Trace("weather updated")
			}
			resp.Body.Close()
		}
		hub.LiveInformationCh <- hub.Message{WsType: hub.Weather, Message: CurrentWeather}
		time.Sleep(interval)
	}
}
