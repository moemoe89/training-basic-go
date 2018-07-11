package controller

import (
	"fmt"
	"net/http"
)

func UrlParam(writer http.ResponseWriter, request *http.Request) {

	// contoh mendapatkan nilai dari url parameter (query) => ?name=asda&address=fkgjflf
	queryValues := request.URL.Query()
	fmt.Fprintf(writer, "url param: %s!\n", queryValues.Get("name"))
	fmt.Fprintf(writer, "url param: %s!\n", queryValues.Get("address"))

}
