package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump() {
	// If this function exits, unregister the client and close the socket
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf("User Disconnected: %v", err)
			break
		}

		c.Hub.broadcast <- message
	}

}

func (c *Client) WritePump() {

	defer c.Conn.Close()

	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Failed to write message: %v", err)
			break
		}
	}
}
