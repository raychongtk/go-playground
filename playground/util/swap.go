package util

func Swap(num1 *int, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}
