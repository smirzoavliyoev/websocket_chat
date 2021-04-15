package main

import (
	"log"
	"net/http"
	"os"
	"websocket_chat/cmd/router"
)

func main() {

	router.CollectRoutes()

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8000"
	}

	log.Println(port)
	// start server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
