package text

// https://leetcode.com/problems/longest-palindrome/
func longestPalindrome(s string) int {
	var length int
	table := make(map[string]int)

	for _, c := range s {
		table[string(c)]++
	}

	for _, v := range table {
		length += v / 2 * 2 // get even number from a value, e.g 21 will be 20
	}

	if length < len(s) {
		length++
	}
	return length
}
