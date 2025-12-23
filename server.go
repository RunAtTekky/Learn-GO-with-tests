package goapp

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")

	switch player {
	case "RunAt":
		fmt.Fprint(writer, "20")
	case "Messi":
		fmt.Fprint(writer, "8")
	}
}
