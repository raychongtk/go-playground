package math

func FindMax(numbers []int) int {
	return findMax(numbers)
}

func findMax(numbers []int) int {
	max := 0

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max
}
