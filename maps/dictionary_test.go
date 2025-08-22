package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"naruto": "A shonen anime."}

	t.Run("Known Test", func(t *testing.T) {

		got, _ := dictionary.Search("naruto")
		want := "A shonen anime."

		assertString(t, got, want)
	})

	t.Run("Unknown Test", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNoEntryFound

		if err == nil {
			t.Fatal("Wanted an error but got none")
		}

		assertError(t, err, want)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
