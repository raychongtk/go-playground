package text

// https://leetcode.com/problems/find-common-characters/
func commonChars(A []string) []string {
	table := [26]int{}

	for _, c := range A[0] {
		table[c-'a']++
	}

	for i := 1; i < len(A); i++ {
		table2 := [26]int{}
		for _, c := range A[i] {
			table2[c-'a']++
		}

		for j := 0; j < 26; j++ {
			if table2[j] < table[j] {
				table[j] = table2[j]
			}
		}
	}

	var result []string
	for i := 0; i < 26; i++ {
		for j := 0; j < table[i]; j++ {
			result = append(result, string(rune(i+'a')))
		}
	}

	return result
}
