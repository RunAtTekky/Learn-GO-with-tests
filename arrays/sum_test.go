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

	check_sums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d but want %d", got, want)
		}
	}

	t.Run("Test sum of tails of slices on empty slice", func(t *testing.T) {
		got := SumTails([]int{1, 2, 3, 4}, []int{1, 2, 3}, []int{1}, []int{})
		want := []int{9, 5, 0, 0}

		check_sums(t, got, want)
	})

	t.Run("Test sum of tails of slices", func(t *testing.T) {
		got := SumTails([]int{1, 2, 3, 4}, []int{1, 2, 3}, []int{1})
		want := []int{9, 5, 0}

		check_sums(t, got, want)
	})
}

func TestBadBank(t *testing.T) {
	const (
		RealMadrid = `Real Madrid`
		PSG        = `PSG`
		Barcelona  = `Barcelona`
		Juventus   = `Juventus`
		Chelsea    = `Chelsea`
	)
	transactions := []Transaction{
		{
			From: PSG,
			To:   Barcelona,
			Sum:  222,
		},
		{
			From: Juventus,
			To:   RealMadrid,
			Sum:  100,
		},
		{
			From: RealMadrid,
			To:   Chelsea,
			Sum:  120,
		},
	}

	balance_PSG := BalanceFor(transactions, PSG)
	balance_Real_Madrid := BalanceFor(transactions, RealMadrid)
	balance_Barcelona := BalanceFor(transactions, Barcelona)
	balance_Juventus := BalanceFor(transactions, Juventus)
	balance_Chelsea := BalanceFor(transactions, Chelsea)

	AssertEqual(t, balance_PSG, -222)
	AssertEqual(t, balance_Real_Madrid, -20)
	AssertEqual(t, balance_Barcelona, 222)
	AssertEqual(t, balance_Juventus, -100)
	AssertEqual(t, balance_Chelsea, 120)
}

func TestReduce(t *testing.T) {
	t.Run("Test multiplying of all elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		multiply := func(res, x int) int {
			return res * x
		}

		got := Reduce(numbers, multiply, 1)
		want := 120

		AssertEqual(t, got, want)
	})
	t.Run("Test concatenation of all strings", func(t *testing.T) {
		names := []string{"neymar", "messi", "ronaldo"}

		concat := func(res, x string) string {
			return res + x
		}

		got := Reduce(names, concat, "")
		want := "neymarmessironaldo"

		AssertEqual(t, got, want)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Fatalf("Did not want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Fatalf("got %v but want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Fatalf("got %v but want false", got)
	}
}
