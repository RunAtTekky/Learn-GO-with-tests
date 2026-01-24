package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		svr := NewPlayerServer(&StubPlayerStore{})

		request := newGameRequest()
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		AssertStatus(t, response, http.StatusOK)
	})
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func TestLeague(t *testing.T) {
	t.Run("Returns 200 on /league", func(t *testing.T) {
		wantedLeague := League{
			{"RunAt", 10},
			{"Ronaldo", 5},
			// {"Messi", 8},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		svr := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		AssertStatus(t, response, http.StatusOK)
		AssertContentType(t, response, jsonContentType)
		AssertLeague(t, got, wantedLeague)
	})
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response is not of type %s, got %v", want, response.Result().Header)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) League {
	t.Helper()
	league, _ := NewLeague(body)
	return league
}

func AssertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
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

		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("Return Messi's Score", func(t *testing.T) {
		request := newGetScoreRequest("Messi")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "8")
	})

	t.Run("Return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Benzema")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		AssertStatus(t, response, http.StatusNotFound)
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
		AssertStatus(t, response, http.StatusAccepted)

		AssertPlayerWin(t, &store, player)
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
