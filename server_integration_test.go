package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()

	store := &FileSystemPlayerStore{database}
	svr := NewPlayerServer(store)
	player := "RunAt"

	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest(player)

		svr.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"RunAt", 3},
		}

		assertLeague(t, got, want)
	})

}
