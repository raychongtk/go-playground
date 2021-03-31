package text

// https://leetcode.com/problems/maximum-number-of-balloons/
func maxNumberOfBalloons(text string) int {
	var count int
	table := make([]int, 26)
	balloon := "balloon"

	for _, c := range text {
		table[c-'a']++
	}

	for true {
		for _, b := range balloon {
			table[b-'a']--

			if table[b-'a'] < 0 {
				return count
			}
		}
		count++
	}

	return count
}
