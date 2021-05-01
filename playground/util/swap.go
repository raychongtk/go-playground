package util

func Swap(num1 *int, num2 *int) {
	*num1, *num2 = *num2, *num1
}
