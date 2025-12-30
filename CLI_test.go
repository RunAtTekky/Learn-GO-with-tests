package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/runattekky/go-app"
)

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
		in := strings.NewReader("RunAt wins\n")
		playerStore := &poker.StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "RunAt")
	})

	t.Run("record Ronaldo win from user input", func(t *testing.T) {
		in := strings.NewReader("Ronaldo wins\n")
		playerStore := &poker.StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Ronaldo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Mbappe wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, blindAlerter)
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
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	if got.amount != want.amount {
		t.Errorf("got amount %d but want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, but want %v", got.at, want.at)
	}
}
