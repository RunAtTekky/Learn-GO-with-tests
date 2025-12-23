package main

import (
	goapp "github.com/runattekky/go-app"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(goapp.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
