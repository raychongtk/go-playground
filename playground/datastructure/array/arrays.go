package array

func SameArray(numbers1 []int, numbers2 []int) bool {
	if len(numbers1) != len(numbers2) {
		return false
	}

	for i := 0; i < len(numbers1); i++ {
		if numbers1[i] != numbers2[i] {
			return false
		}
	}

	return true
}

func ReverseArray(numbers []int) {
	start := 0
	end := len(numbers) - 1

	for start < end {
		numbers[start], numbers[end] = numbers[end], numbers[start]
		start++
		end--
	}
}

// https://leetcode.com/problems/sort-array-by-parity/
func SortArrayByParity(numbers []int) {
	var index int

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			numbers[index], numbers[i] = numbers[i], numbers[index]
			index++
		}
	}
}
