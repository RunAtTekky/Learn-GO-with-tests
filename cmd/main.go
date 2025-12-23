package main

import (
	"log"
	"net/http"

	"github.com/runattekky/go-app/server"
)

func main() {
	svr := &server.PlayerServer{Store: NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", svr))
}
