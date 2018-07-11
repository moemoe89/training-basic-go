package controller

import (
	"fmt"
	"net/http"
	"strconv"
)

type luas struct {
	Panjang int
	Lebar   int
}

// contoh membuat fungsi based on struct
func(l luas) Result() int {
	return l.Panjang * l.Lebar
}

func HelloWorld(writer http.ResponseWriter, request *http.Request) {

	var l luas
	l.Panjang = 10
	l.Lebar = 5

	// contoh mengakses fungsi
	result := l.Result()

	fmt.Fprint(writer, "hello world: ", strconv.Itoa(result))
}
