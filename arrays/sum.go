package arrays

func Sum(numbers [5]int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}
