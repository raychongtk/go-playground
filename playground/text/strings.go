package text

import "strings"

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

func Join(elements []string, separator string) string {
	var builder strings.Builder

	for _, element := range elements {
		if builder.Len() > 0 {
			builder.WriteString(separator)
		}
		builder.WriteString(element)
	}

	return builder.String()
}

func IsPalindrome(text string) bool {
	start := 0
	end := len(text) - 1
	chars := []rune(text)

	for start < end {
		if chars[start] != chars[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func ReverseString(text string) string {
	var builder strings.Builder

	for i := len(text) - 1; i >= 0; i-- {
		builder.WriteString(string(text[i]))
	}

	return builder.String()
}
