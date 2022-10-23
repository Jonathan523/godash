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

type launchpadInformation struct {
	Title     string
	Host      string
	Bookmarks []bookmark.Bookmark
	Weather   weather.OpenWeatherApiResponse
	System    system.System
}

func launchpad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	files.ParseAndServeHtml(w, "index.gohtml", launchpadInformation{
		Title:     "Godash",
		Bookmarks: bookmark.Bookmarks,
		Weather:   weather.CurrentOpenWeather,
		System:    system.Sys,
	})
}

// @Schemes
// @Summary     get the current weather
// @Description gets the current weather
// @Tags        weather
// @Produce     json
// @Success     200 {object} weather.OpenWeatherApiResponse
// @Success     204 {object} message.Response
// @Router      /weather [get]
func getWeather(w http.ResponseWriter, r *http.Request) {
	if weather.Conf.OpenWeather.Key != "" {
		jsonResponse(w, weather.CurrentOpenWeather, http.StatusOK)
	} else {
		jsonResponse(w, message.Response{Message: message.NotFound.String()}, http.StatusNoContent)
	}
}

// @Schemes
// @Summary     live system information
// @Description gets live information of the system
// @Tags        system
// @Produce     json
// @Success     200 {object} system.LiveInformation
// @Success     204 {object} message.Response
// @Router      /system/live [get]
func routeLiveSystem(w http.ResponseWriter, r *http.Request) {
	if system.Config.LiveSystem {
		jsonResponse(w, system.Sys.Live, http.StatusOK)
	} else {
		jsonResponse(w, message.Response{Message: message.NotFound.String()}, http.StatusNoContent)
	}
}

// @Schemes
// @Summary     static system information
// @Description gets static information of the system
// @Tags        system
// @Produce     json
// @Success     200 {object} system.StaticInformation
// @Success     204 {object} message.Response
// @Router      /system/static [get]
func routeStaticSystem(w http.ResponseWriter, r *http.Request) {
	if system.Config.LiveSystem {
		jsonResponse(w, system.Sys.Static, http.StatusOK)
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
