package router

import (
	"net/http"
	"websocket_chat/cmd/handlers"
)

func CollectRoutes() {
	http.Handle("/", handlers.WebFileServer())
	http.HandleFunc("/ws", handlers.WebSocketChats)
}
