package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("Error getting store from file %s, %v", db.Name(), err)
	}

	svr := NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", svr); err != nil {
		log.Fatalf("Could not listen on PORT :5000 %v", err)
	}
}
