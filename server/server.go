package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		p.processWin(writer)
	case http.MethodGet:
		p.showScore(writer, request)
	}
}

func (p *PlayerServer) processWin(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	score := p.Store.GetPlayerScore(player)
	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, score)

}
