package main

import (
	"WelcomeGO/myutil"
	"fmt"
    "WelcomeGO/LearnPrintStatement"
)

func main() {
	fmt.Println("Welcome to Learning GoLang !!")
	myutil.Helper("This message from Helper func")
    LearnPrintStatement.LearnPrint();
	//variables 
	var name string = "Himanshu"
	fmt.Println(name)
	 // shortcut use 
	persion := "ShortCut Himanshu"
	fmt.Println(persion)

	// var currency int = 10 
	var currency = 10
	fmt.Println(currency)

	// var decimalCurr = 10.5  // not required to aasign type of variables, it decides at compile time
	var decimalCurr float32 = 100.5 // decide variable type compile time
	fmt.Println(decimalCurr)

	const pi = 22.7
	//pi = 23.7 const value could not be reassign
	fmt.Println(pi)

	// we can Re-assign var type  variables
	var chanse bool = false
	chanse = true // can be change because of var type
	fmt.Println(chanse)

	// Exported from other file, SO there variable name should be staring from Capital letters
	fmt.Println(myutil.Persion)

}
