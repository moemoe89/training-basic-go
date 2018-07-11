package controller

import (
	"fmt"
	"net/http"
)

func InputPost(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {

		// contoh mendapatkan nilai dari input body form-data dan x-www-form-urlencoded
		name := request.FormValue("name")
		address := request.FormValue("address")
		fmt.Fprintf(writer, "Name = %s\n", name)
		fmt.Fprintf(writer, "Address = %s\n", address)

		return
	}

	http.Error(writer, "Invalid request method.", http.StatusMethodNotAllowed)
}
