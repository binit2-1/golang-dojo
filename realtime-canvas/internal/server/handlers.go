package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	ws "realtime-canvas/internal/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &ws.Client{
		Hub: hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	log.Println("New WebSocket Connection Established!")
	defer conn.Close()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

}
