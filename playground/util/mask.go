package util

func Mask(text string) string {
	if len(text) == 0 {
		return text
	} else if len(text) == 1 {
		return "******"
	}

	return string(text[0]) + "******" + string(text[len(text)-1])
}

func MaskLeft(text string) string {
	if len(text) == 0 {
		return text
	} else if len(text) == 1 {
		return "******"
	}

	chars := []rune(text)

	for i := 0; i < len(chars)-4; i++ {
		chars[i] = '*'
	}

	return string(chars)
}
