package controller

import (
	"fmt"
	"net/http"
)

func UrlPath(writer http.ResponseWriter, request *http.Request) {

	// contoh untuk mendapatkan value dari suatu path
	id := request.URL.Path[len("/url-path/"):]
	fmt.Fprintf(writer, "path: %s!\n", id)

}
