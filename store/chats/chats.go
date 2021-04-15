package chats

import "websocket_chat/store/message"

var (
	broadcast = make(chan message.Message)
)
