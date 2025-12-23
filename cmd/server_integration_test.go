package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/runattekky/go-app/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{}
	svr := server.PlayerServer{Store: &store}
	player := "RunAt"

	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	request := newGetScoreRequest(player)

	svr.ServeHTTP(response, request)

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
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
