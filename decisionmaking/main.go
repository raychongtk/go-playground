package main

import "fmt"

func main() {
	var mark int
	var grade string
	mark = 80

	switch mark {
	case 100:
		grade = "A"
	case 90:
		grade = "B"
	case 60, 70, 80:
		grade = "C"
	case 50:
		grade = "D"
	default:
		grade = "E"
	}

	switch {
	case grade == "A":
		fmt.Println("Excellent")
	case grade == "B", grade == "C":
		fmt.Println("Good")
	case grade == "D":
		fmt.Println("You passed")
	case grade == "E":
		fmt.Println("You failed")
	}

	fmt.Println("your grade is", grade)
}
