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

func ReverseArray(numbers1 []int) {
	start := 0
	end := len(numbers1) - 1

	for start < end {
		temp := numbers1[start]
		numbers1[start] = numbers1[end]
		numbers1[end] = temp
		start++
		end--
	}
}
