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
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "runat"
		definition := "golang developer"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("Existing word", func(t *testing.T) {
		word := "runat"
		definition := "golang developer"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("Updating already present word test", func(t *testing.T) {
		word := "go"
		old_definition := "fav programming language"

		dict := Dictionary{word: old_definition}

		new_definition := "favourite programming language"
		err := dict.Update(word, new_definition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, new_definition)
	})

	t.Run("Updating not present word test", func(t *testing.T) {
		dict := Dictionary{}
		word := "go"
		new_definition := "favourite programming language"

		err := dict.Update(word, new_definition)

		assertError(t, err, ErrCannotUpdateWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete existing word", func(t *testing.T) {
		word := "runat"
		definition := "golang developer"

		dict := Dictionary{word: definition}

		err := dict.Delete(word)

		assertError(t, err, nil)
	})

	t.Run("Delete non-existing word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("runat")

		assertError(t, err, ErrCannotDeleteWordDoesNotExist)
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

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("Should find added word: ", err)
	}

	assertString(t, got, definition)
}
