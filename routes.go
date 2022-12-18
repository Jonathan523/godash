package main

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"godash/hub"
)

var (
	upgrader = websocket.Upgrader{}
)

func (g *goDash) ws(c echo.Context) error {
	if g.config.PageUrl.String() != c.Request().Header.Get("Origin") {
		return errors.New("bad request")
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	messageChan := make(hub.NotifierChan)
	g.hub.NewClients <- messageChan
	defer func() {
		g.hub.ClosingClients <- messageChan
		ws.Close()
	}()

	go func() {
		defer ws.Close()
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				break
			}
		}
	}()

	for {
		select {
		case msg, ok := <-messageChan:
			if !ok {
				err := ws.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					return err
				}
				return err
			}
			err := ws.WriteJSON(msg)
			if err != nil {
				return err
			}
		}
	}

}
