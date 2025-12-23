package goapp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestGETPlayers(t *testing.T) {
	t.Run("Return RunAt's Score", func(t *testing.T) {
		request := newGetScoreRequest("RunAt")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("Return Messi's Score", func(t *testing.T) {
		request := newGetScoreRequest("Messi")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "8")
	})
}
