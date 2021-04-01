package math

func fib(n int) int {
	if n == 1 {
		return 1
	}

	var result int
	first := 0
	second := 1

	for i := 0; i < n-1; i++ {
		result = first + second
		first = second
		second = result
	}

	return result
}
