package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
    number_int := 123 // Declare and initialize
    number_int = 456  // Update the value of already initialized variable 
	num = 24 // num should be already declared 
	*/

	fmt.Println("Welcome to dataConversion !")
	var digit int = 25
	fmt.Println("value of digit is :" ,digit)
	fmt.Printf("Datatype of digit is :%T\n", digit)

	// if we try to add float into digit then it will error
	//digit += 2.5 // error before conversion of digit into float
	var data float32 = float32(digit)  // we assign 
	fmt.Println("value of data is :" ,data)
	// now we can add float into digit as is type of float
	data += 2.5
	fmt.Println("now value of data is :", data)
	/*
	In Go, you cannot directly change the data type of a variable once it has been declared. 
	This is because Go is a statically typed language,
    meaning the type of a variable is determined at the time of its declaration and cannot be changed later.
	*/
	fmt.Printf("Datatype of digit is :%T\n", digit) //we can't change datatype of digit as it was type of int at time of declaration

	number_int := 123
	number_string := strconv.Itoa(number_int) // Itoa --> converts Integer to alphabates(string)

	fmt.Println("value of number_string is :" , number_string)

	fmt.Printf("Datatype of number_string is :%T\n", number_string)

	number_string = "1234"
	number_integer , _:= strconv.Atoi(number_string) // Atoi --> alphabates to integer conv , it returns (value , error)

	fmt.Println("value of number_int is :" , number_integer)

	fmt.Printf("Datatype of number_int is :%T\n", number_integer)

	


}
