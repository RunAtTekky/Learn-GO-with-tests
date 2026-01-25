package poker

import (
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func AssertStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Result().StatusCode != want {
		t.Errorf("Did not get correct STATUS, got %d but want %d", got.Result().StatusCode, want)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to record win but want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("got %s calls to record win but want %s", store.WinCalls[0], winner)
	}
}
