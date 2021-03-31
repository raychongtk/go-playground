package text

// merge string alternately
func MergeString(str1 string, str2 string) string {
	var result string
	var maxLen int
	len1 := len(str1)
	len2 := len(str2)

	if len1 > len2 {
		maxLen = len1
	} else {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result += string(str1[i])
		}
		if i < len2 {
			result += string(str2[i])
		}
	}

	return result
}
