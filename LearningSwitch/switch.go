package main

import "fmt"

func main() {
	var day int
	fmt.Println("Please enter number between 1 - 7")
	fmt.Scan(&day) // user input will store into day 
	switch day {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
		// handling multiple cases into sigle flow
	case 4,5,6:
		fmt.Println("Today is your taken leave days")
    default: // in case of 7 we are handling in default
       fmt.Println("Enjoy ! Today is Sunday")
	}

	var month string
	fmt.Println("please enter months name")
	fmt.Scan(&month)

	switch month{
		
	case "January" ,"February" ,"March": // any month out of these will be winter 
	fmt.Println("Now it is Winter season")
	case "April" , "May" ,"June":
		fmt.Println("It is summer season")
	case "October","November", "December":
		fmt.Println("It is Cold time")
	}

	// switch statement also works with expression
	temperature := 20
	switch{
	case  temperature < 20:
		fmt.Println("It is cold outside")
	case temperature >= 20:
		fmt.Println("It is Hot outside")
	case temperature <= 0:
		fmt.Println("Freezing outside")
	}
}
