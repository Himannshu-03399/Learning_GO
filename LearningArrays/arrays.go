package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Learning arrays in go !")
	// array declaration Syntax: { var nameofArray [size_Array] data_types}
	var friends [4]string  // declare array
	//puttting values in array names friends
	friends[0] = "Himanshu" // assign vlaues in it
	friends[1] = "Ayush"
	friends[2] = "vinay"
    friends[3] = "Anurag"

   // printing values of array
    fmt.Println("These are my friends :",friends) 
	// this time declare and assign values same time
	var prices = [5] int{10,20,30,40,50}
	fmt.Println("prices of product is :" , prices)

	// know lenth of array
	fmt.Println("size of prices array is :",len(prices))
	//access values of arrays
	fmt.Println("value of prices at 2nd index is :" , prices[2])

	// default values in arrays of different diff datatypes
	var intArr [5] float32 
	fmt.Println("default values in intArr are :" , intArr) // both for int & float default values are [0 0 0 ...]

	var StrArr [3] string // default values "" 
	fmt.Printf("vlaues in StrArr are %q\n" , StrArr) // string outout in "quoted-string" using printf func
	// arrays of default bool value
	var boolAr [3] bool
	fmt.Println("default values in boolArr is :", boolAr)

	// Array of pointers
    var pointerArray [2]*int
    fmt.Println(pointerArray) // Output: [<nil> <nil>]

    // Array of complex numbers
    var complexArray [2]complex128
    fmt.Println(complexArray) // Output: [(0+0i) (0+0i)]
}
