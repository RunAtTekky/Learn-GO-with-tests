package goapp

import (
	"fmt"
	"net/http"
)

func PlayerServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "20")
}
