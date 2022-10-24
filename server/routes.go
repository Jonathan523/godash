package server

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"godash/bookmark"
	"godash/files"
	"godash/hub"
	"godash/system"
	"godash/weather"
	"net/http"
)

type launchpadInformation struct {
	Title     string
	Host      string
	Bookmarks []bookmark.Bookmark
	Weather   weather.Weather
	System    system.System
}

func launchpad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	files.ParseAndServeHtml(w, "index.gohtml", launchpadInformation{
		Title:     "GoDash",
		Bookmarks: bookmark.Bookmarks,
		Weather:   weather.CurrentWeather,
		System:    system.Sys,
	})
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithField("error", err).Warning("Cannot upgrade websocket")
		return
	}
	messageChan := make(hub.NotifierChan)
	server.Hub.NewClients <- messageChan
	defer func() {
		server.Hub.ClosingClients <- messageChan
		conn.Close()
	}()
	go readPump(conn)
	for {
		select {
		case msg, ok := <-messageChan:
			if !ok {
				err := conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					return
				}
				return
			}
			err := conn.WriteJSON(msg)
			if err != nil {
				return
			}
		}
	}
}
