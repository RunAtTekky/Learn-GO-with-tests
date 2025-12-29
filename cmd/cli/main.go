package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/runattekky/go-app"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type '{Name} wins' to record win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening file %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	cli := poker.NewCLI(store, os.Stdin)
	cli.PlayPoker()
}
