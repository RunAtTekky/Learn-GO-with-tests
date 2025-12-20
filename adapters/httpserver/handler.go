package httpserver

import (
	"fmt"
	"net/http"

	"github.com/runattekky/go-specs-greet/domain/interactions"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, interactions.Greet(name))
}
