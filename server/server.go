package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		p.processWin(writer, request)
	case http.MethodGet:
		p.showScore(writer, request)
	}
}

func (p *PlayerServer) processWin(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusAccepted)
	player := getPlayer(request.URL.Path)
	p.Store.RecordWin(player)
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, request *http.Request) {
	player := getPlayer(request.URL.Path)
	score := p.Store.GetPlayerScore(player)
	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, score)

}

func getPlayer(path string) string {
	player := strings.TrimPrefix(path, "/players/")
	return player
}
