package handlers

import (
	"net/http"
)

func WebFileServer() http.Handler {

	return http.FileServer(http.Dir("./handlers/public"))
}
