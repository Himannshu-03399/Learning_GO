package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey ! body , What's your Name")

	// 1st method
	// var name string // define variable to store user input
	// fmt.Scan(&name) //whatever user will type via keyboard, it stores in addresss of name till system not gets whitespace
	// Now printout 
	// fmt.Println("Hello my code viewer , Mr.", name)  // output: Hello my code viewer , Mr. {given input by user}

   // 2nd Standered way to take complete user input through (os.stdin)
    reader := bufio.NewReader(os.Stdin) // buffer input/output is temp storage
	
	Name, _ := reader.ReadString('\n') // it store user input in Name variable from temp storage reader.('\n') defines take input(till next line)
    fmt.Println("Hey Strainger :)",Name)
	
}