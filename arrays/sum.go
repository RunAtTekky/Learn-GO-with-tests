package arrays

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	balance := 0.0
	for _, transaction := range transactions {
		if transaction.From == name {
			balance -= transaction.Sum
		} else if transaction.To == name {
			balance += transaction.Sum
		}
	}
	return balance
}

func Reduce[T any](collection []T, combining_function func(T, T) T, initialVal T) T {
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
