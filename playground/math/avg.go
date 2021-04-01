package math

import "math"

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
func AvgWithoutMinMax(numbers []int) float64 {
	min := math.MaxInt32
	max := 0
	sum := 0

	for _, num := range numbers {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}

		sum += num
	}

	return float64(sum-min-max) / float64(len(numbers)-2)
}
