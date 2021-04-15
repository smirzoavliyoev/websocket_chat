package handlers

import (
	"log"
	"net/http"
	"websocket_chat/utils"
)

func WebFileServer() http.Handler {

	path, err := utils.GetDirOrFilePathFromRoot(".", "public")
	if err != nil {
		panic(err)
	}
	log.Println(path)

	return http.FileServer(http.Dir(path))
}
