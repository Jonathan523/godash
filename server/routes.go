package server

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"godash/bookmark"
	"godash/files"
	"godash/hub"
	"godash/message"
	"godash/system"
	"godash/weather"
	"net/http"
)

type LaunchpadInformation struct {
	Title     string
	Host      string
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

func routeLiveSystem(w http.ResponseWriter, r *http.Request) {
	if system.Config.LiveSystem {
		jsonResponse(w, system.Live.System.Live, http.StatusOK)
	} else {
		jsonResponse(w, message.Response{Message: message.NotFound.String()}, http.StatusNoContent)
	}
}

func routeStaticSystem(w http.ResponseWriter, r *http.Request) {
	if system.Config.LiveSystem {
		jsonResponse(w, system.Live.System.Static, http.StatusOK)
	} else {
		jsonResponse(w, message.Response{Message: message.NotFound.String()}, http.StatusNoContent)
	}
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithField("error", err).Warning("Cannot upgrade websocket")
		return
	}
	messageChan := make(hub.NotifierChan)
	system.Live.Hub.NewClients <- messageChan
	defer func() {
		system.Live.Hub.ClosingClients <- messageChan
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
