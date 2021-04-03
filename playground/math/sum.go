package math

// https://leetcode.com/problems/running-sum-of-1d-array/
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

func SumRecursion(numbers []int) int {
	return sumRecursion(numbers, 0)
}

func sumRecursion(numbers []int, idx int) int {
	if idx == len(numbers) {
		return 0
	}

	return numbers[idx] + sumRecursion(numbers, idx+1)
}

func AccumulateSum(target int) int {
	if target == 1 {
		return 1
	}

	return target + AccumulateSum(target-1)
}
