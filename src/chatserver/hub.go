package chatserver


type Hub struct {

	Clients map[*Client]bool

	Broadcast chan []byte

	Register chan *Client

	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <- h.Register:
			h.Clients[client] = true
			for hclient := range h.Clients {
				select {
				case hclient.Send <- []byte("login"):
				default:
					close(hclient.Send)
					delete(h.Clients, hclient)
				}
			}
		case client := <- h.Unregister:
			if _, ok := h.Clients[client]; ok {
				for hclient := range h.Clients {
					select {
					case hclient.Send <- []byte("logout"):
					default:
						close(hclient.Send)
						delete(h.Clients, hclient)
					}
				}
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <- h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
