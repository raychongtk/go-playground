package dictionary

func singleNumber(nums []int) int {
	dictionary := make(map[int]bool)

	for _, value := range nums {
		if _, ok := dictionary[value]; ok {
			dictionary[value] = true
		} else {
			dictionary[value] = false
		}
	}

	for key, value := range dictionary {
		if !value {
			return key
		}
	}

	return -1
}
