package arrays

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(curr_balance float64, t Transaction) float64 {
		if t.From == name {
			return curr_balance - t.Sum
		}

		if t.To == name {
			return curr_balance + t.Sum
		}

		return curr_balance
	}

	return Reduce(transactions, adjustBalance, 0.0)
}

func Reduce[A, B any](collection []A, combining_function func(B, A) B, initialVal B) B {
	result := initialVal
	for _, val := range collection {
		result = combining_function(result, val)
	}
	return result
}

// Returns sum of a slice
func Sum(numbers []int) int {
	adder := func(res, x int) int { return res + x }
	return Reduce(numbers, adder, 0)
}

// Returns a slice with sum of all slices
func SumAll(numbers_to_sum ...[]int) []int {
	allAdder := func(res, x []int) []int {
		return append(res, Sum(x))
	}

	return Reduce(numbers_to_sum, allAdder, []int{})
}

// Returns a slice with sum of all slices except first element of each slice
func SumTails(tails_to_sum ...[]int) (res []int) {
	tailAdder := func(res, x []int) []int {
		if len(x) == 0 {
			return append(res, 0)
		} else {
			tail := x[1:]
			return append(res, Sum(tail))
		}
	}

	return Reduce(tails_to_sum, tailAdder, []int{})
}
