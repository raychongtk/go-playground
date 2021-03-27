package math

func FindMax(numbers []int) int {
	return max(numbers)
}

func max(numbers []int) int {
	max := 0

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max
}
