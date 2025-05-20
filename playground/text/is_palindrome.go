package text

func isPalindrome(x int) bool {
	var reverseInteger int
	var actualNumber = x

	for actualNumber > 0 {
		num := actualNumber % 10
		actualNumber /= 10
		reverseInteger = reverseInteger*10 + num
	}

	return x == reverseInteger
}
