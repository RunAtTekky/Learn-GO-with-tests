package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://google.com"
	fastURL := "https://blog.runat.xyz"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
