package server

import (
	"launchpad/bookmark"
	"launchpad/files"
	"net/http"
)

type LaunchpadInformation struct {
	Title     string
	Bookmarks []bookmark.Bookmark
}

func launchpad(w http.ResponseWriter, r *http.Request) {
	files.ParseHtml(w, "index.gohtml", LaunchpadInformation{Title: "Launchpad", Bookmarks: bookmark.Bookmarks})
}

func serveStyleCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/style.css")
}
