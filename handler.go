package go_specs_greet

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %s", request.URL.Query().Get("name"))
}
