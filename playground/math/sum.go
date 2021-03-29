package math

func RunningSum(numbers []int) []int {
	runningSum := make([]int, len(numbers))

	runningSum[0] = numbers[0]
	for i := 1; i < len(numbers); i++ {
		runningSum[i] = runningSum[i-1] + numbers[i]
	}

	return runningSum
}

func Sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
