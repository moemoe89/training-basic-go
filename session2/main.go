package main

import (
	"fmt"
	"net/http"

	// import dapat di alias sehingga cukup dituliskan ctrl daripada controller
	ctrl "training-basic-go/session2/controller"
)


func main(){

	// endpoint dengan direct function
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "root path")
	})

	// contoh endpoint dengan function terpisah
	http.HandleFunc("/hello-world", ctrl.HelloWorld)
	http.HandleFunc("/test-post", ctrl.TestPost)
	http.HandleFunc("/url-param", ctrl.UrlParam)
	http.HandleFunc("/url-path/", ctrl.UrlPath)
	http.HandleFunc("/input-post", ctrl.InputPost)
	http.HandleFunc("/input-json", ctrl.InputJSON)

	// start server
	http.ListenAndServe(":8080", nil)
}
