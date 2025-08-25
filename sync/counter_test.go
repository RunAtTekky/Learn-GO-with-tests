package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times and leave at it", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		got := counter.Value()
		want := 3

		assertCounter(t, got, want)
	})

	t.Run("counter runs safely concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for range wantCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		got := counter.Value()

		assertCounter(t, got, wantCount)
	})
}

func assertCounter(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
