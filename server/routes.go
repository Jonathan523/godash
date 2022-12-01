package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/websocket"
	"github.com/sirupsen/logrus"
	"godash/bookmark"
	"godash/hub"
	"godash/system"
	"godash/weather"
)

func (server *Server) goDash(c context.Context, ctx *app.RequestContext) {
	ctx.HTML(consts.StatusOK, "index.gohtml", utils.H{
		"Title":   server.Title,
		"Entries": bookmark.Entries,
		"Weather": weather.CurrentWeather,
		"System":  system.Sys,
	})
}

func webSocket(_ context.Context, ctx *app.RequestContext) {
	err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
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
	})
	if err != nil {
		logrus.WithField("error", err).Warning("Cannot upgrade websocket")
		return
	}
}
