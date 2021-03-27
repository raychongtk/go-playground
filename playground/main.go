package main

import (
	"fmt"
	myMath "playground/math"
)

func main() {
	numbers := []int{5, 2, 1, 4, 5, 6, 8, 9, 0, 10}
	min := myMath.FindMin(numbers)
	max := myMath.FindMax(numbers)
	fmt.Printf("min=%d, max=%d\n", min, max)
}
