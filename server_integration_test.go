package poker_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	poker "github.com/runattekky/go-app"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := poker.NewFileSystemPlayerStore(database.(*os.File))
	AssertNoError(t, err)

	svr, err := poker.NewPlayerServer(store)
	if err != nil {
		t.Fatalf("Problem creating player server %v", err)
	}

	player := "RunAt"

	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest(player)

		svr.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := poker.League{
			{"RunAt", 3},
		}

		AssertLeague(t, got, want)
	})

}
