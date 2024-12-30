package main

import "fmt"

func divide(a, b float64) (res float64, err error) {

	// denominator shouldn't be 0
	if b == 0 {
		return 0 ,fmt.Errorf("denominator must not be Zero") // for safe we are returning 0 instead of infinity 
	}
	res = a/b;
	return res , nil // in this case only res will return as error will be nil

}

// note we can also return string instead of error type
func division( x float32 , y float32) (float32,string){
	if y == 0 {
		return 0 , "Denominator shouldn't be zero"
	}
	// var res = x/y
	return x/y , "No error cautch"
	
}

func main(){
	// continue with happly flow
	fmt.Println("Here we are learning Error handling in GO !")
	ans, _ := divide(10,5) // after ans variable "_" is positions of error message handling we can write as (ans,err)
	fmt.Println("It is ans of Function without error, Result is :", ans)

	// this time we will handle error
	fmt.Println("Now we are Handling error")
    res , err := divide(10,0)
	//this time error is there so err variable has something to give 
	if err != nil{
		fmt.Println("Error handling  occures because of denominator is 0")
	}
	fmt.Println("Division of given numbers is,Returning 0 instead of infinity " , res)

	// Using String type errorHandling 
	fmt.Println("USing String type error ")
	value ,str := division(10,5)
	fmt.Println("division of two numbers is:" ,value , str ) // print value with No error string 

	//in case of error string
	val , str := division(10,0)
	fmt.Println("division of two numbers is :" , val ,str) // in case of string type error 

}
