package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) Run() {

	for {

		select {
		case cli := <-hub.register:
			hub.onConnect(cli)

		case cli := <-hub.unregister:
			hub.onDisconnects(cli)
		}
	}

}

func (h *Hub) BroadCast(msg interface{}, ignore *Client) {
	data, _ := json.Marshal(msg)
	for _, cli := range h.clients {
		if cli != ignore {
			cli.outband <- data
		}
	}
}

func (hub *Hub) onConnect(c *Client) {
	log.Println(" -> Client Connected! ", c.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	c.id = c.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, c)
}

func (hub *Hub) onDisconnects(cli *Client) {
	log.Println(" -> Client Disconect! ", cli.socket.RemoteAddr())

	cli.Close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	i := -1
	for j, c := range hub.clients {
		if c.id == cli.id {
			i = j
			break
		}
	}

	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]

}

func (h *Hub) WsHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cl := NewClient(h, socket)
	h.register <- cl

	go cl.Write()
}
