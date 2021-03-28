package main

import "fmt"

func main() {
	// integer
	var age int
	age = 100
	fmt.Println("I am", age, "years old")

	// boolean & string
	var status bool
	var statusText string
	status = true
	if status {
		statusText = "on"
	} else {
		statusText = "off"
	}

	fmt.Println("light is turned", statusText)

	// floating point
	var pi float32
	var pi2 float64
	pi = 3.1415926535897932384626433832795028841971693993
	pi2 = 3.1415926535897932384626433832795028841971693993
	fmt.Println(pi, pi2)
}
