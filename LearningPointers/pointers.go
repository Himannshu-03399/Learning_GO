package main

import "fmt"

func modifyValueByReference(num *int){ // passing num pointer to this func
    *num = *num + 20 // 5 + 20 = 25 
}

func main() {
	fmt.Println("Welcome to learning pointers ! ,pointer is also variable which stores memory address of another variable")
	//delclare variable and pointer
	num := 2
	name:= "Himanshu" 

	// var ptr* int //int represents memory address where data  kept is type of int
	ptr := &num
	// var ptr1* string
	ptr1 := & name // we cann't keep address of name as it is type of string and we define pointer of int type 
	// ptr1 = &name // ptr1 is type of string so we can keep address of string type data memory address


	fmt.Println("vlaue of num is :" , num)
	fmt.Println("memory address of num is :",ptr)

	fmt.Println("vlaue of name is :" , name)
	fmt.Println("memory address  of name is :",ptr1) // this prints address of name variable

	// accessing actual value which is store in memory address
	fmt.Println("value of name which is stored at memory loacation ptr1 is : " , *ptr1)

	var ptr2 *int // default value of ptr2 is nill now as we didn't store memory address of any int variable in ptr2
	if ptr2 == nil{
		fmt.Println("ptr2 is not assigned yet !")
	}

	value := 5
	modifyValueByReference(&value)//passing address of vlaue variable 
	fmt.Println("Value after Modify by reference is :" , value)



}