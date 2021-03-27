package main

import "fmt"

func main() {
	// for loop
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// while loop
	loopCount := 0

	for loopCount < 10 {
		loopCount++
	}

	fmt.Println(loopCount)

	// for each range
	stringArray := [10]string{"hello", "world"}
	for index, str := range stringArray {
		fmt.Println(index, str)
	}
}
