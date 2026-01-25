package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/runattekky/go-app"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Problem opening the file %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("Problem creating file system from store, %v", err)
	}

	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)

	svr, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("Problem creating player server %v", err)
	}

	log.Fatal(http.ListenAndServe(":5000", svr))
}
