package main

import "fmt"

func main() {
	fmt.Println("Welcome to Learning functions in GO")

	//add:=sum(20,10) this type of variable declaration only valid within functions , otherwise ( var ans = sum(10,20) in general)
	add := sum(10,5) // storing output of sum function into add variable
	fmt.Println("Addition of two digits using sum is :" ,add)
	product := multiply(5.9 , 8)
	fmt.Println("Multiplication of x and y is" , product)

}
// sum(a int , b) this means we give datatypes to only a 
func sum(a,b int ) (result int){ // we can also write sun( a int , b int )
	result = a+b
	return 
} 

func multiply(x float32 , y float32) (res float32){
	res = x*y;
    return
}
