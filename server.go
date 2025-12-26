package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"RunAt", 10},
		{"Ronaldo", 5},
		{"Messi", 8},
	}
}

func (p *PlayerServer) playersHandler(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	switch request.Method {
	case http.MethodPost:
		p.processWin(writer, player)
	case http.MethodGet:
		p.showScore(writer, player)
	}
}

func (p *PlayerServer) processWin(writer http.ResponseWriter, player string) {
	writer.WriteHeader(http.StatusAccepted)
	p.store.RecordWin(player)
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, score)
}
