package arrays

func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}

func SumAll(numbers_to_sum ...[]int) []int {
	n := len(numbers_to_sum)
	res := make([]int, n)

	for i, numbers := range numbers_to_sum {
		res[i] = Sum(numbers)
	}
	return res
}

func SumTails(tails_to_sum ...[]int) (res []int) {
	for _, numbers := range tails_to_sum {
		res = append(res, Sum(numbers[1:]))
	}
	return
}
