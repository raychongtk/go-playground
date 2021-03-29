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
