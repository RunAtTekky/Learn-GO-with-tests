package poker

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

		store, err := NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)
		got := store.GetLeague()

		want := League{
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

		store, err := NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)

		got := store.GetPlayerScore("RunAt")
		want := 10
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "RunAt", "Wins": 10},
		{"Name": "Ronaldo", "Wins": 5}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)

		store.RecordWin("RunAt")

		got := store.GetPlayerScore("RunAt")
		want := 11
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "RunAt", "Wins": 10},
		{"Name": "Ronaldo", "Wins": 5}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)

		store.RecordWin("Messi")

		got := store.GetPlayerScore("Messi")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)
	})

	t.Run("return league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Ronaldo", "Wins": 5},
			{"Name": "RunAt", "Wins": 10}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{Name: "RunAt", Wins: 10},
			{Name: "Ronaldo", Wins: 5},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
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

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
