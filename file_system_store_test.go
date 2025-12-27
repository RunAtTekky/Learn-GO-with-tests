package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get league", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name": "RunAt", "Wins": 10},
		{"Name": "Ronaldo", "Wins": 5}]`)

		store := FileSystemPlayerStore{database}
		got := store.GetLeague()

		want := []Player{
			{"RunAt", 10},
			{"Ronaldo", 5},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name": "RunAt", "Wins": 10},
		{"Name": "Ronaldo", "Wins": 5}]`)
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("RunAt")
		want := 10
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
