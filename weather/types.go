package weather

type PackageConfig struct {
	Location    Location
	OpenWeather OpenWeather `mapstructure:"OPEN_WEATHER"`
	Current     OpenWeatherApiResponse
}

type Location struct {
	Latitude  float32
	Longitude float32
}

type OpenWeather struct {
	Key    string
	Units  string
	Lang   string
	Digits bool
}

type Weather struct {
	Icon        string  `json:"icon"`
	Temp        float64 `json:"temp"`
	Description string  `json:"description"`
	Humidity    uint8   `json:"humidity"`
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
	Temp     float64 `json:"temp"`
	Humidity uint8   `json:"humidity"`
}

type OpenWeatherApiSys struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}
