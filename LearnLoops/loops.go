package main

import "fmt"

/*
In go only for loop is there , while and doWhile loops are not in golang
*/

func main() {
	fmt.Println("welcome to loops in go")
	for i := 0; i < 5; i++ {
		fmt.Println("Number is :", i)
	}

	// while loop can be written in form of for loop
	count := 1
	for {
		fmt.Println("Now count is :", count)
		count++
		if count == 5 { // break when we want to terminate
			break
		}
	}

	// Use of range keyword to iterate over slice / map / arrays
	marks:= []int {10,20,25,5,40}
	for roll, obtainMarks := range marks{
		fmt.Printf("Roll number of Student is %d , and Marks for this Student is : %d \n", roll, obtainMarks)
	}
}
