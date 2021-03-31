package text

import (
	"strings"
)

/*
	https://leetcode.com/problems/word-pattern/
	Input: pattern = "abba", s = "dog cat cat dog"
	Output: true
*/
func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	set := make(map[string]int)
	asciiTable := [26]string{}

	for i, c := range pattern {
		if _, ok := set[words[i]]; ok && asciiTable[c-'a'] == "" {
			return false
		}

		if asciiTable[c-'a'] != "" && asciiTable[c-'a'] != words[i] {
			return false
		}

		asciiTable[c-'a'] = words[i]
		set[words[i]] = 1
	}

	return true
}
