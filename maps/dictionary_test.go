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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "runat"
	definition := "golang developer"

	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
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

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("Should find added word: ", err)
	}

	assertString(t, got, definition)
}
