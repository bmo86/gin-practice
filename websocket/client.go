package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	hub     *Hub
	id      int
	socket  *websocket.Conn
	outband chan []byte
}

func NewClient(h *Hub, s *websocket.Conn) *Client {
	return &Client{
		hub:     h,
		socket:  s,
		outband: make(chan []byte),
	}
}

func (c *Client) Write() {
	for {
		select {
		case message, ok := <-c.outband:

			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c Client) Close() {
	c.socket.Close()
	close(c.outband)
}
