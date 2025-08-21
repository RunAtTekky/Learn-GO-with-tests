package arrays

func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}
