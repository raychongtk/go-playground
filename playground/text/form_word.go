package text

// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
func countCharacters(words []string, chars string) int {
	var count = 0
	table := make([]int, 26)

	for _, c := range chars {
		table[c-'a']++
	}

	for _, word := range words {
		table2 := make([]int, 26)
		canBeFormed := true

		for _, c := range word {
			table2[c-'a']++

			if table2[c-'a'] > table[c-'a'] {
				canBeFormed = false
				break
			}
		}

		if canBeFormed {
			count += len(word)
		}
	}

	return count
}
