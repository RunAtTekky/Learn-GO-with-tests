package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("RunAt wins\n")
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("Expected a win call but did not get any")
	}

	got := playerStore.winCalls[0]
	want := "RunAt"

	if got != want {
		t.Errorf("did not record correct winner, got %s but want %s", got, want)
	}
}
