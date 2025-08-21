package arrays

// Returns sum of a slice
func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}

// Returns a slice with sum of all slices
func SumAll(numbers_to_sum ...[]int) []int {
	n := len(numbers_to_sum)
	res := make([]int, n)

	for i, numbers := range numbers_to_sum {
		res[i] = Sum(numbers)
	}
	return res
}

// Returns a slice with sum of all slices except first element of each slice
func SumTails(tails_to_sum ...[]int) (res []int) {
	for _, numbers := range tails_to_sum {
		if len(numbers) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(numbers[1:]))
		}
	}
	return
}
