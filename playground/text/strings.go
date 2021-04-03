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

func IsContainingDuplicateCharacter(text string) bool {
	asciiTable := make([]int, 26)

	for _, c := range text {
		if asciiTable[c-'a'] == 1 {
			return true
		}
		asciiTable[c-'a'] += 1
	}
	return false
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

func ReverseStringRecursion(text string) string {
	return string(reverseStringRecursion([]rune(text), 0, len(text)-1))
}

func reverseStringRecursion(text []rune, idx1 int, idx2 int) []rune {
	if idx1 == idx2 {
		return text
	}

	text[idx1], text[idx2] = text[idx2], text[idx1]
	return reverseStringRecursion(text, idx1+1, idx2-1)
}

func ReverseWord(text string) string {
	builder := new(strings.Builder)
	parts := strings.Split(text, " ")

	for i := len(parts) - 1; i >= 0; i-- {
		if builder.Len() > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(parts[i])
	}

	return builder.String()
}
