package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get league", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "RunAt", "Wins": 10},
		{"Name": "Ronaldo", "Wins": 5}]`)
		defer cleanDatabase()

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
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "RunAt", "Wins": 10},
		{"Name": "Ronaldo", "Wins": 5}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("RunAt")
		want := 10
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
