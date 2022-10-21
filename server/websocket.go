package server

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func inAllowedHosts(str string) bool {
	for _, a := range server.AllowedHosts {
		if a == str {
			return true
		}
	}
	return false
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return inAllowedHosts(r.Header.Get("Origin"))
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func readPump(conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
