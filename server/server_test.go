package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/runattekky/go-app/server"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"RunAt":     20,
			"Cristiano": 5,
			"Messi":     8,
		},
	}

	svr := &server.PlayerServer{Store: &store}

	t.Run("Return RunAt's Score", func(t *testing.T) {
		request := newGetScoreRequest("RunAt")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("Return Messi's Score", func(t *testing.T) {
		request := newGetScoreRequest("Messi")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "8")
	})

	t.Run("Return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Benzema")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
