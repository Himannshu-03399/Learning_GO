package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps")
	// declare map name as marks using make func
	marks := make(map[string]int)
	//entering marks of students along with name
	marks["Alice"] = 45
	fmt.Println("Marks of Alice before appending is :", marks["Alice"]) // output 45
	marks["Bob"] = 70
	marks["Vinay"] = 65
	marks["Himanshu"] = 80
	marks["Ayush"] = 85
	fmt.Println("Marks of Vinay is :", marks["Vinay"])

	// map in go return {vlaue of key , bool} for every enrty, if entry will not there in map then boolian will be false and value for that query wll be 0
	//deleting entry of Vinay
	delete(marks, "Vinay")
	// after deleting vinay marks will be 0 as that entry is no longer
	value, isPresent := marks["vinay"]
	fmt.Println("After delete Marks of Vinay is :", value)
	fmt.Println("Is Vinay Entry is now exist in map :", isPresent)

	//we can also change value of entry
	marks["Alice"] = 50                                                // overwirte
	fmt.Println("Marks of Alice after overright is :", marks["Alice"]) // output 50

	//using range keyword 
	for student , number := range marks{
		fmt.Printf("marks for %s and obtainedMarks is :% d\n" , student , number)
	}
    // declare map and assign vlaue same time
	BookPrice := map[string]int{
		"Atomic Habits" :100,
		"How to enfluence" : 85,
		"Godan PremChand" : 120,
	}

	// printing price of books
	for book , price := range BookPrice{
		fmt.Printf("Name of Book is %s and Price of this book is : %d\n" , book , price)
	}

}
