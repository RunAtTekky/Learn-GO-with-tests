package go_specs_greet

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello, World")
}
