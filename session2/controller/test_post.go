package controller

import (
	"fmt"
	"net/http"
)

func TestPost(writer http.ResponseWriter, request *http.Request) {

	// pengecekan method yang digunakan client ketika dihit
	if request.Method == "POST" {
		fmt.Fprint(writer, "Ini request post")
		return
	}

	// http lib punya bank code untuk menampilkan status code
	http.Error(writer, "Invalid request method.", http.StatusMethodNotAllowed)
}
