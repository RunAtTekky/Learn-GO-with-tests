package httpserver

import (
	"fmt"
	"net/http"

	"github.com/runattekky/go-specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, replyWith(interactions.Curse))
	return mux
}

func replyWith(f func(name string) string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("name")
		fmt.Fprint(writer, f(name))
	}

}
