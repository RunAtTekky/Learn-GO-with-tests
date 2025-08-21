package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum of slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d but want %d given %v", got, want, numbers)
		}
	})
}
