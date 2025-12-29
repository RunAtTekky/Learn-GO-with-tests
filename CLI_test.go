package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record RunAt win from user input", func(t *testing.T) {
		in := strings.NewReader("RunAt wins\n")
		playerStore := &StubPlayerStore{}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		if len(playerStore.winCalls) != 1 {
			t.Fatal("Expected a win call but did not get any")
		}

		AssertPlayerWin(t, playerStore, "RunAt")
	})

	t.Run("record Ronaldo win from user input", func(t *testing.T) {
		in := strings.NewReader("Ronaldo wins\n")
		playerStore := &StubPlayerStore{}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		if len(playerStore.winCalls) != 1 {
			t.Fatal("Expected a win call but did not get any")
		}

		AssertPlayerWin(t, playerStore, "Ronaldo")
	})
}
