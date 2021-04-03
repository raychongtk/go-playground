package math

// https://leetcode.com/problems/maximum-number-of-balls-in-a-box/
func countBalls(lowLimit int, highLimit int) int {
	boxes := make([]int, highLimit+1)
	maxBall := 0
	curIdx := 0

	for i := lowLimit; i <= highLimit; i++ {
		curIdx = numSum(i)
		boxes[curIdx]++
		maxBall = max(maxBall, boxes[curIdx])
	}

	return maxBall
}

func numSum(num int) int {
	var sum int

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
