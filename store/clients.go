package store

import "github.com/gorilla/websocket"

var Clients = make(map[*websocket.Conn]bool)
