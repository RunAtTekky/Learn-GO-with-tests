package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/runattekky/go-specs-greet"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, go_specs_greet.Greet(name))
}
