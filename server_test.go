package goapp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Return RunAt's Score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/RunAt", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Return Messi's Score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Messi", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "8"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
