package goapp

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")

	fmt.Fprint(writer, GetPlayerScore(player))

}

func GetPlayerScore(player string) string {
	switch player {
	case "RunAt":
		return "20"
	case "Messi":
		return "8"
	}
	return ""
}
