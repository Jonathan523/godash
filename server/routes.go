package server

import (
	"launchpad/bookmark"
	"launchpad/files"
	"launchpad/weather"
	"net/http"
)

type LaunchpadInformation struct {
	Title     string
	Bookmarks []bookmark.Bookmark
	Weather   weather.OpenWeatherApiResponse
}

func launchpad(w http.ResponseWriter, r *http.Request) {
	files.ParseHtml(w, "index.gohtml", LaunchpadInformation{
		Title:     "Launchpad",
		Bookmarks: bookmark.Bookmarks,
		Weather:   weather.CurrentOpenWeather,
	})
}
