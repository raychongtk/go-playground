package math

import "math"

func FindMin(numbers []int) int {
	return min(numbers)
}

func min(numbers []int) int {
	min := math.MaxInt32

	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	return min
}
