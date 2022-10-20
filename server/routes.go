package server

import (
	"godash/bookmark"
	"godash/files"
	"godash/weather"
	"net/http"
)

type LaunchpadInformation struct {
	Title     string
	Bookmarks []bookmark.Bookmark
	Weather   weather.OpenWeatherApiResponse
}

func launchpad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	files.ParseAndServeHtml(w, "index.gohtml", LaunchpadInformation{
		Title:     "Godash",
		Bookmarks: bookmark.Bookmarks,
		Weather:   weather.CurrentOpenWeather,
	})
}
