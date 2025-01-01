package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Learning Conditions in GO !")
	x := 5
	fmt.Println("Hello user please enter number")
	var digit int
	fmt.Scan(&digit)

	if x > digit {
		fmt.Println("Yes x is greter than user input")
	} else if digit == x {
		fmt.Println("User input is equal to x")
	} else {
		fmt.Println("Input is greater than x")
	}
}
