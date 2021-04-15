package handlers

import (
	"log"
	"net/http"
	"websocket_chat/services/channel"
	"websocket_chat/store"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		EnableCompression: true,
		ReadBufferSize:    1 << 10,
		WriteBufferSize:   1 << 10,
	}
)

func WebSocketChats(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("errr", err)
	}
	channel := channel.NewChannel(ws)

	defer ws.Close()

	store.Clients[ws] = true

	defer channel.Conn.Close()

	for {
		channel.Reader()
	}
}
