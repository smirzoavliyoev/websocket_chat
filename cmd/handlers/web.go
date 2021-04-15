package handlers

import (
	"log"
	"net/http"
	"os"
)

func WebFileServer() http.Handler {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dir)

	return http.FileServer(http.Dir(dir + "/public"))
}
