package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat v 5 times", func(t *testing.T) {
		repeated := Repeat("v", 5)
		expected := "vvvvv"

		if repeated != expected {
			t.Errorf("got %q but want %q", repeated, expected)
		}
	})

	t.Run("Repeat 'RunAt' 2 times", func(t *testing.T) {
		repeated := Repeat("RunAt", 2)
		expected := "RunAtRunAt"

		if repeated != expected {
			t.Errorf("got %q but want %q", repeated, expected)
		}

	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("v", 2)
	}
}
