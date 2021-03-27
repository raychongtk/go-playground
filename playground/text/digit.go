package text

func IsNumber(text string) bool {
	for _, ascii := range text {
		if !isDigit(ascii) {
			return false
		}
	}
	return true
}

func isDigit(ascii int32) bool {
	return ascii >= 48 && ascii <= 57
}
