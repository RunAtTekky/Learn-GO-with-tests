package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	svr := PlayerServer{Store: NewInMemoryPlayerStore()}
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
