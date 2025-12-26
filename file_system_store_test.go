package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
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
}
