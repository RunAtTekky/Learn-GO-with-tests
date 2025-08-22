package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := map[string]string{"naruto": "A shonen anime."}

	got := Search(dictionary, "naruto")
	want := "A shonen anime."

	if got != want {
		t.Errorf("got %q but want %q given, %q", got, want, dictionary)
	}
}
