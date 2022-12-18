package hub

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
		Notifier          NotifierChan
		NewClients        chan NotifierChan
		ClosingClients    chan NotifierChan
		clients           map[NotifierChan]struct{}
		LiveInformationCh chan Message
	}
)

func NewHub() *Hub {
	hub := Hub{}
	hub.LiveInformationCh = make(chan Message)
	hub.Notifier = make(NotifierChan)
	hub.NewClients = make(chan NotifierChan)
	hub.ClosingClients = make(chan NotifierChan)
	hub.clients = make(map[NotifierChan]struct{})
	go hub.listen()
	go func() {
		for {
			if msg, ok := <-hub.LiveInformationCh; ok {
				hub.Notifier <- msg
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
		case s := <-h.ClosingClients:
			delete(h.clients, s)
		case event := <-h.Notifier:
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
