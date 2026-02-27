package websocket


type Hub struct{
	clients map[*Client]bool
	Broadcast chan []byte
	Register chan *Client
	Unregister chan *Client
}

func NewHub() *Hub{
	return &Hub{
		clients: make(map[*Client]bool),
		Broadcast: make(chan []byte),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run(){
	for{
		select{
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.Unregister:
			delete(h.clients, client)
		case message := <-h.Broadcast:
			for client := range h.clients{
				client.Send <- message
			}

		}
	}
}