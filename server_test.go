package poker_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	poker "github.com/runattekky/go-app"
)

var (
	dummyGame = &GameSpy{}
	tenMS     = 10 * time.Millisecond
)

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		svr := mustMakePlayerServer(t, &poker.StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})

	t.Run("start game with 3 players and declare RunAt as winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		winner := "RunAt"

		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		svr := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		wsURL := "ws" + strings.TrimPrefix(svr.URL, "http") + "/ws"
		ws := mustDialWS(t, wsURL)

		defer svr.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(tenMS)
		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, winner)

		within(t, tenMS, func() { assertWebSocketGotMsg(t, ws, wantedBlindAlert) })
	})
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func TestLeague(t *testing.T) {
	t.Run("Returns 200 on /league", func(t *testing.T) {
		wantedLeague := poker.League{
			{"RunAt", 10},
			{"Ronaldo", 5},
			// {"Messi", 8},
		}

		store := poker.StubPlayerStore{nil, nil, wantedLeague}
		svr := mustMakePlayerServer(t, &store, dummyGame)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response, http.StatusOK)
		AssertContentType(t, response, "application/json")
		AssertLeague(t, got, wantedLeague)
	})
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response is not of type %s, got %v", want, response.Result().Header)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) poker.League {
	t.Helper()
	league, _ := poker.NewLeague(body)
	return league
}

func AssertLeague(t testing.TB, got, want poker.League) {
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
	store := poker.StubPlayerStore{
		map[string]int{
			"RunAt":     20,
			"Cristiano": 5,
			"Messi":     8,
		},
		[]string{},
		nil,
	}

	svr := mustMakePlayerServer(t, &store, dummyGame)

	t.Run("Return RunAt's Score", func(t *testing.T) {
		request := newGetScoreRequest("RunAt")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("Return Messi's Score", func(t *testing.T) {
		request := newGetScoreRequest("Messi")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "8")
	})

	t.Run("Return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Benzema")
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	svr := mustMakePlayerServer(t, &store, dummyGame)
	t.Run("records win when POST", func(t *testing.T) {
		player := "Messi"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusAccepted)

		poker.AssertPlayerWin(t, &store, player)
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

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game poker.Game) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store, dummyGame)
	if err != nil {
		t.Fatalf("Problem creating player server %v", err)
	}

	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Could not open a WS connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("Could not send message over WS connection %v", err)
	}
}

func assertStatus(t *testing.T, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Code != want {
		t.Errorf("did not get correct status code, got %d but want %d", got.Code, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong, got %q but want %q", got, want)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()
	done := make(chan struct{}, 1)
	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("Timed out")
	case <-done:
	}
}

func assertWebSocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf("got %q but want %q", string(msg), want)
	}

}
