package text

func IsEmpty(text string) bool {
	return len(text) == 0
}

func TrimLeft(text string) string {
	var result string
	flag := false

	for _, t := range text {
		if !flag && t == 32 {
			continue
		}

		flag = true
		result += string(t)
	}

	return result
}

func TrimRight(text string) string {
	var result string
	flag := false

	for i := len(text) - 1; i > 0; i-- {
		if !flag && text[i] == 32 {
			continue
		}

		flag = true
		result = string(text[i]) + result
	}

	return result
}

func RemoveDuplicateCharacter(text string) string {
	table := make(map[string]int)
	var result string

	for _, t := range text {
		key := string(t)
		if table[key] == 0 {
			result += key
		}

		table[key] = 1
	}

	return result
}
