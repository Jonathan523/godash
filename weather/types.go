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

type Weather struct {
	Icon        string  `json:"icon"`
	Temp        float32 `json:"temp"`
	Description string  `json:"description"`
	Humidity    float32 `json:"humidity"`
	Sunrise     string  `json:"sunrise"`
	Sunset      string  `json:"sunset"`
	Units       string  `json:"units"`
}

type OpenWeatherApiResponse struct {
	Weather []OpenWeatherApiWeather `json:"weather"`
	Main    OpenWeatherApiMain      `json:"main"`
	Sys     OpenWeatherApiSys       `json:"sys"`
}

type OpenWeatherApiWeather struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type OpenWeatherApiMain struct {
	Temp     float32 `json:"temp"`
	Humidity float32 `json:"humidity"`
}

type OpenWeatherApiSys struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}
