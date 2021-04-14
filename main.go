package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Channel struct {
	conn *websocket.Conn
	send chan Message
}

func (this *Channel) reader() {

	var (
		msg Message
	)

	for {
		err := this.conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			delete(clients, this.conn)
		}

		this.send <- msg
	}
}

func (this *Channel) writer() {

	for {
		msg := <-this.send
		for client := range clients {

			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				delete(clients, this.conn)
			}
		}
	}
}

func NewChannel(conn *websocket.Conn) Channel {
	channel := Channel{
		conn: conn,
		send: make(chan Message, 0),
	}
	// go channel.reader()
	go channel.writer()

	return channel
}

func AnotherOne(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("errr", err)
	}
	channel := NewChannel(ws)

	defer ws.Close()

	clients[ws] = true

	defer channel.conn.Close()

	for {
		channel.reader()
	}
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
)

var (
	upgrader = websocket.Upgrader{
		EnableCompression: true,
		ReadBufferSize:    1 << 10,
		WriteBufferSize:   1 << 10,
	}
)

type Message struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/ws", AnotherOne)

	log.Println("OK")

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	log.Println(port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
