package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times and leave at it", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3

		assertCounter(t, counter, want)
	})

	t.Run("counter runs safely concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for range wantCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d but want %d", got.Value(), want)
	}
}
