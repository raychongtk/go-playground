package main

import (
	"fmt"
	myMath "playground/math"
	myText "playground/text"
	myUtil "playground/util"
)

func main() {
	numbers := []int{5, 2, 1, 4, 5, 6, 8, 9, 0, 10}
	min := myMath.FindMin(numbers)
	max := myMath.FindMax(numbers)
	fmt.Printf("min=%d, max=%d\n", min, max)

	text := "012345678999995"
	text2 := "       0123456789abcd       "
	fmt.Println(myText.IsNumber(text))
	fmt.Println(myText.IsNumber(text2))

	var str string
	str = "aaaa"
	fmt.Println(myText.IsEmpty(str))

	// trim spaces
	fmt.Println(myText.TrimLeft(text2))
	fmt.Println(myText.TrimRight(text2))

	// remove duplicated char
	fmt.Println(myText.RemoveDuplicateCharacter(text))

	// swapping value
	num1 := 10
	num2 := 20
	myUtil.Swap(&num1, &num2)
	fmt.Println(num1, num2)

	// join string
	var elements []string
	elements = append(elements, "test", "test1", "test2", "test3", "test4", "test5", "test6", "test7")
	fmt.Println(myText.Join(elements, ", "))
}
