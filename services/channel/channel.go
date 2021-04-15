package channel

import (
	"log"
	"websocket_chat/store"
	"websocket_chat/store/message"

	"github.com/gorilla/websocket"
)

type Channel struct {
	Conn *websocket.Conn
	Send chan message.Message
}

func (this *Channel) Reader() {

	var (
		msg message.Message
	)

	for {
		err := this.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			delete(store.Clients, this.Conn)
		}

		this.Send <- msg
	}
}

func (this *Channel) Writer() {

	for {
		msg := <-this.Send
		for client := range store.Clients {

			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				delete(store.Clients, this.Conn)
			}
		}
	}
}

func NewChannel(conn *websocket.Conn) Channel {
	channel := Channel{
		Conn: conn,
		Send: make(chan message.Message, 0),
	}
	// go channel.reader()
	go channel.Writer()

	return channel
}
