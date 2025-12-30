package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/runattekky/go-app"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s *scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

func TestCLI(t *testing.T) {
	t.Run("record RunAt win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nRunAt wins\n")
		playerStore := &poker.StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummyStdOut, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "RunAt")
	})

	t.Run("record Ronaldo win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nRonaldo wins\n")
		playerStore := &poker.StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummyStdOut, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Ronaldo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nMbappe wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummyStdOut, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", want.amount, want.at), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was NOT scheduled for %v", want.amount, want.at)
				}

				gotAlert := blindAlerter.alerts[i]
				assertScheduledAlert(t, gotAlert, want)
			})
		}
	})

	t.Run("it prompts the user to enter number of players", func(t *testing.T) {
		in := strings.NewReader("7\n")
		stdout := &bytes.Buffer{}
		blindAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(dummyPlayerStore, in, stdout, blindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(want.String(), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was NOT scheduled for %v", i, blindAlerter.alerts)
				}

				gotAlert := blindAlerter.alerts[i]
				assertScheduledAlert(t, gotAlert, want)
			})
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	if got.amount != want.amount {
		t.Errorf("got amount %d but want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, but want %v", got.at, want.at)
	}
}
