package math

func TopTwoElements(numbers []int) (number1 int, number2 int) {
	var max1, max2 int

	for _, num := range numbers {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}

	return max1, max2
}

func TopThreeElements(numbers []int) (number1, number2, number3 int) {
	var max1, max2, max3 int

	for _, num := range numbers {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}
	}

	return max1, max2, max3
}
