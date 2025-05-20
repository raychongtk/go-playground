package text

func romanToInt(s string) int {
	var number int
	i := 0
	for i < len(s) {
		if i+1 < len(s) {
			specialRoman := s[i : i+2]
			specialNumber := convertSpecialRoman(specialRoman)
			if specialNumber == -1 {
				number += convertRoman(string(s[i]))
			} else {
				number += specialNumber
				i++
			}
		} else {
			number += convertRoman(string(s[i]))
		}
		i++
	}
	return number
}

func convertRoman(s string) int {
	if s == "I" {
		return 1
	} else if s == "V" {
		return 5
	} else if s == "X" {
		return 10
	} else if s == "L" {
		return 50
	} else if s == "C" {
		return 100
	} else if s == "D" {
		return 500
	} else if s == "M" {
		return 1000
	} else {
		return -1
	}
}

func convertSpecialRoman(s string) int {
	if s == "IV" {
		return 4
	} else if s == "IX" {
		return 9
	} else if s == "XL" {
		return 40
	} else if s == "XC" {
		return 90
	} else if s == "CD" {
		return 400
	} else if s == "CM" {
		return 900
	} else {
		return -1
	}
}
