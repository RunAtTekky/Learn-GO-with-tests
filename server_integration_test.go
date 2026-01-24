package poker

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database.(*os.File))
	AssertNoError(t, err)

	svr := NewPlayerServer(store)
	player := "RunAt"

	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest(player)

		svr.ServeHTTP(response, request)

		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, newLeagueRequest())
		AssertStatus(t, response, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := League{
			{"RunAt", 3},
		}

		AssertLeague(t, got, want)
	})

}
