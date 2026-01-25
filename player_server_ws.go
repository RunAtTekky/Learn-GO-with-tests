package poker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type playerServerWS struct {
	*websocket.Conn
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection to websockets %v\n", err)
	}

	return &playerServerWS{conn}
}

func (w *playerServerWS) waitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("Error reading from websocket %v\n", err)
	}

	return string(msg)
}

func (w *playerServerWS) Write(p []byte) (n int, err error) {
	err = w.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}

	return len(p), err
}
