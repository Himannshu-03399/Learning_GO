package main

import "fmt"

func main() {
	fmt.Println("Welcome to learning Slice , a efficient & dynamic size data structure !")
	// initialize and store at same time
	numbers :=[] int {} // syntax
	fmt.Println("numbers stores " , numbers)
	/*
	SLICES having len , capacity (capacity every times gets double)
	*/
	fmt.Println("default size of numberes is :" , len(numbers) )
	fmt.Println("deafult capacity  of numberes is : " , cap(numbers))

	// append  something to slices here it is numberes
	numbers = append(numbers,1,2,3,4)
	fmt.Println("numbers became after inserting : " ,numbers) 

	// printing length and capacity after inserting 
	fmt.Println("size of numberes after inserting values is :" , len(numbers) ) // output 4 
	fmt.Println("now capacity of numberes is : " , cap(numbers)) // output 4 
	// now if append 1 more number
	numbers = append(numbers, 6)
	fmt.Println("now capacity of numberes is : " , cap(numbers)) // capacity gets double so this time output 8 

	// create slices using make func : len 3 & capacity 5
	marks := make([]int , 3 , 5)
	//printing slices info/
	fmt.Println("slices is : " , marks) // default value will be 0
	fmt.Println("Length is " , len(marks) , "capacity is" , cap(marks))

	// after appending 
	marks = append(marks, 30,40,50)
	fmt.Println("slices is : " , marks)
	fmt.Println("Length is " , len(marks) , "capacity is" , cap(marks)) // gets double output 10


	
}