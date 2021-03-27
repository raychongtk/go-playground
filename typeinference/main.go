package main

import "fmt"

func main() {
	// type inference
	text := "this is a text"
	fmt.Println(text)
	fmt.Printf("the type of text is %T\n", text)

	// mixed type inference
	a, b, c := 1, 2.1, "c"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("the type of a is %T\n", a)
	fmt.Printf("the type of b is %T\n", b)
	fmt.Printf("the type of c is %T\n", c)
}
