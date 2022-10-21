package weather

type Config struct {
	Location    Location
	OpenWeather OpenWeather `mapstructure:"OPEN_WEATHER"`
	Current     OpenWeatherApiResponse
}

type Location struct {
	Latitude  float32
	Longitude float32
}

type OpenWeather struct {
	Key   string
	Units string
	Lang  string
}

type OpenWeatherApiResponse struct {
	Weather []OpenWeatherApiWeather `json:"weather" `
	Main    OpenWeatherApiMain      `json:"main" `
	Sys     OpenWeatherApiSys       `json:"sys" `
	Name    string                  `json:"name" `
	Units   string                  `json:"units" `
}

type OpenWeatherApiWeather struct {
	Description string `json:"description" `
	Icon        string `json:"icon" `
}

type OpenWeatherApiMain struct {
	Temp     float32 `json:"temp" `
	Humidity float32 `json:"humidity" `
}

type OpenWeatherApiSys struct {
	Sunrise    int64  `json:"sunrise"`
	Sunset     int64  `json:"sunset"`
	StrSunrise string `json:"str_sunrise"`
	StrSunset  string `json:"str_sunset"`
}
