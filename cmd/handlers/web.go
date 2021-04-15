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
	return http.FileServer(http.Dir(dir + "/public"))
}
