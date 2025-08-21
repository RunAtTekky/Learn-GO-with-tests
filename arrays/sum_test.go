package arrays

import (
	"reflect"
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

func TestSumAll(t *testing.T) {
	t.Run("Test sum of all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{3, 9})
		want := []int{6, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d but want %d", got, want)
		}
	})

	t.Run("Test sum of 3 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{3, 9}, []int{1})
		want := []int{6, 12, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}

func TestSumTails(t *testing.T) {
	t.Run("Test sum of tails of slices", func(t *testing.T) {
		got := SumTails([]int{1, 2, 3, 4}, []int{1, 2, 3}, []int{1})
		want := []int{9, 5, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}
