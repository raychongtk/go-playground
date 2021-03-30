package math

func SplitInteger(num int) []int {
	var result []int

	for num > 0 {
		remainder := num % 10
		result = append([]int{remainder}, result...)
		num /= 10
	}

	return result
}
