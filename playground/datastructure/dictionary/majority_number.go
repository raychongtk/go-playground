package dictionary

func majorityElement(nums []int) int {
	majority := len(nums) / 2
	m := make(map[int]int)

	for _, num := range nums {
		m[num] = m[num] + 1

		if value, ok := m[num]; ok {
			if value > majority {
				return num
			}
		}
	}

	return -1
}
