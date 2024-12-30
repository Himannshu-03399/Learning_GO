package main

import (
	"WelcomeGO/myutil"
	"fmt"
    "WelcomeGO/LearnPrintStatement"
	
)

/*
If we have multiple Go files with the same package name (main or whatever)
and each defines its own main() function within the same folder,
the Go compiler will throw an error when you attempt to build or run entire package
or more than one file at the same time. However, if you run a single file at a time,
only that file and its dependencies are considered, and no error or warning occurs
because other main() functions in the folder are ignored.

*/

func main() {
	fmt.Println("Welcome to Learning GoLang !!")
	myutil.Helper("This message from Helper func")
	myutil.Demo("Note, if you want to export variable or methods outside of package/folder then it should be start with UpperCase ")
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
	myutil.Helper("Just now myutils/Helper function is called")


}
