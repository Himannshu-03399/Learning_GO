package LearnPrintStatement

import "fmt"

var Name string = "Himashu Raj"
var age float32 = 22.5

const Address string = "India"

func LearnPrint() {

	fmt.Printf("Name of this code author is : %s\n ", Name) // same like c programing printf , println is use

	fmt.Printf("Age of Author is : %.2f \n", age) // after decimal till 2 digits float value will be print
	// print all variable in signle line
	fmt.Printf("Auther Name is : %s , age is : %.1f and Address is: %s\n", Name, age, Address)

	// find data types of variable
	fmt.Printf("DataType of Address is : %T\n", Address)

}
