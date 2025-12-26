package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestLeague(t *testing.T) {
	t.Run("Returns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"RunAt", 10},
			{"Ronaldo", 5},
			// {"Messi", 8},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		svr := NewPlayerServer(&store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server, %q into slice of Player. %v", response.Body, err)
		}

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v but want %v", got, wantedLeague)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"RunAt":     20,
			"Cristiano": 5,
			"Messi":     8,
		},
		[]string{},
		nil,
	}

	svr := NewPlayerServer(&store)

	t.Run("Return RunAt's Score", func(t *testing.T) {
		request := newGetScoreRequest("RunAt")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("Return Messi's Score", func(t *testing.T) {
		request := newGetScoreRequest("Messi")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "8")
	})

	t.Run("Return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Benzema")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	svr := NewPlayerServer(&store)
	t.Run("records win when POST", func(t *testing.T) {
		player := "Messi"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d but want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("got %s but want %s", store.winCalls[0], player)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Did not get correct STATUS, got %d but want %d", got, want)
	}
}
