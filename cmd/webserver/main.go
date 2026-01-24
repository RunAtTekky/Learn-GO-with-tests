package main

import (
	"log"
	"net/http"

	poker "github.com/runattekky/go-app"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	svr := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", svr); err != nil {
		log.Fatalf("Could not listen on PORT :5000 %v", err)
	}
}
