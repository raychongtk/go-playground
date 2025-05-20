package text

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, str := range strs {
		var newPrefix string
		for i, char := range str {
			if i < len(prefix) && string(char) == string(prefix[i]) {
				newPrefix += string(char)
			} else {
				break
			}
		}
		prefix = newPrefix
	}

	return prefix
}
