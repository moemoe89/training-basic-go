package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func InputJSON(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {

		// membaca request body dan mengubah menjadi byte
		b, err := ioutil.ReadAll(request.Body)
		// defer akan menjalankan line di urutan terakhir dari pembacaan code
		defer request.Body.Close()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal untuk memasukkan nilai byte ke variabel, bisa bertipe map, struct atau tipe data biasa
		var msg Message
		err = json.Unmarshal(b, &msg)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		mapResponse := map[string]interface{}{
			"input_body": msg,
			"custom_response": "output custom",
			"aaaa": "aaaa",
		}

		// membuat content type response berupa json
		writer.Header().Set("Content-Type", "application/json")
		// set status code
		writer.WriteHeader(http.StatusOK)
		// write json
		json.NewEncoder(writer).Encode(mapResponse)

		return
	}

	http.Error(writer, "Invalid request method.", http.StatusMethodNotAllowed)
}
