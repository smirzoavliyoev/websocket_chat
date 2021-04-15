package handlers

import (
	"net/http"
	"websocket_chat/utils"
)

func WebFileServer() http.Handler {

	path, err := utils.GetDirOrFilePathFromRoot(".", "public")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.Dir(path))
}
