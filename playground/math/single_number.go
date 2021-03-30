package math

func SingleNumber(numbers []int) int {
	var result int

	for _, num := range numbers {
		result ^= num
	}

	return result
}
