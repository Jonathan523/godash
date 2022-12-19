package hub

import "go.uber.org/zap"

const (
	Weather WsType = iota
	System
)

type (
	NotifierChan chan Message

	WsType  uint
	Message struct {
		WsType  WsType      `json:"ws_type"`
		Message interface{} `json:"message"`
	}

	Hub struct {
		LiveInformationCh chan Message
		NewClients        chan NotifierChan
		ClosingClients    chan NotifierChan
		logger            *zap.SugaredLogger
		notifier          NotifierChan
		clients           map[NotifierChan]struct{}
	}
)

func NewHub(logger *zap.SugaredLogger) *Hub {
	hub := Hub{
		LiveInformationCh: make(chan Message),
		NewClients:        make(chan NotifierChan),
		ClosingClients:    make(chan NotifierChan),
		logger:            logger,
		notifier:          make(NotifierChan),
		clients:           make(map[NotifierChan]struct{}),
	}
	go hub.listen()
	go func() {
		for {
			if msg, ok := <-hub.LiveInformationCh; ok {
				hub.notifier <- msg
			}
		}
	}()
	return &hub
}

func (h *Hub) listen() {
	for {
		select {
		case s := <-h.NewClients:
			h.clients[s] = struct{}{}
			h.logger.Debugw("websocket connection added", "total clients", len(h.clients))
		case s := <-h.ClosingClients:
			delete(h.clients, s)
			h.logger.Debugw("websocket connection removed", "total clients", len(h.clients))
		case event := <-h.notifier:
			for client := range h.clients {
				select {
				case client <- event:
				default:
					close(client)
					delete(h.clients, client)
				}
			}
		}
	}
}
