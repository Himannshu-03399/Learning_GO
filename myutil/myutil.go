// package main  
package myutil // in single directory all go files should have same package

import "fmt"

var Persion = "Public BLR" // this variable could be export 
var persion1 = "private BLR" // this variable can be access within same package

func Helper(message string){
	fmt.Println(message)
	fmt.Println(persion1) // this is private variable 

}