package datastructure

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
		temp := numbers[start]
		numbers[start] = numbers[end]
		numbers[end] = temp
		start++
		end--
	}
}
