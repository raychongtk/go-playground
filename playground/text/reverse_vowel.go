package text

func reverseVowels(s string) string {
	start := 0
	end := len(s) - 1
	runes := []rune(s)
	for start < end {
		shouldSwapPrefix := isAeiou(runes[start])
		shouldSwapSuffix := isAeiou(runes[end])

		if shouldSwapPrefix && shouldSwapSuffix {
			runes[end], runes[start] = runes[start], runes[end]
			start++
			end--
		} else if !shouldSwapPrefix {
			start++
		} else if !shouldSwapSuffix {
			end--
		} else {
			start++
			end--
		}
	}

	return string(runes)
}

func isAeiou(r rune) bool {
	return r == 'a' || r == 'A' || r == 'e' || r == 'E' || r == 'i' || r == 'I' || r == 'o' || r == 'O' || r == 'u' || r == 'U'
}
