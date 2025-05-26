package text

import "strings"

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

func wordPattern(pattern string, s string) bool {
	patternMap := make(map[byte]string)
	wordMap := make(map[string]byte)

	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	for i := 0; i < len(words); i++ {
		c := pattern[i]
		patternValue, patternExists := patternMap[c]
		wordValue, wordExists := wordMap[words[i]]

		if patternExists && wordExists {
			if patternValue != words[i] || wordValue != c {
				return false
			}
		} else if !patternExists && !wordExists {
			patternMap[c] = words[i]
			wordMap[words[i]] = c
		} else {
			return false
		}
	}

	return true
}
