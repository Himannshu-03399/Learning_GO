package main

import "fmt"

func add(a, b int) (result int) {
	result = a + b
	return
}

func multiple(num1, num2 float32) float32 {
	res := num1 * num2
	return res
}

func main() {
	fmt.Println("welcome to defer keyword , it works in LIFO order ")

	// defer keyword is used to execute that line of code at last in Last in First out(LIFO)
	Addition := add(5,7)
	defer fmt.Println("Addition of two digit is :" , Addition) // although it is  written before multiplication but due to defer it will execute at last

	multiplication := multiple(2.5,6)
	fmt.Println("Multiplication is :" , multiplication)

}
