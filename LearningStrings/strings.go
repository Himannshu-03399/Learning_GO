package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to learning string package ")
	data := "Himanshu , Ayush , Vinay , Patel"
	// seperate data based on comma
	name := strings.Split(data, ",")
	fmt.Println("names of friends are :" , name)

	words := "hello world hello Mohan Hello Himanshu"
	cnt := strings.Count(words , "hello")

	fmt.Println("counts of hello in words are :",cnt)

	sentence := "   Hey ! world      "
	trimedSentence := strings.TrimSpace(sentence)
	fmt.Println("sentence converted into trimedSentence is :" , trimedSentence)

	firstName := "Himanshu"
	lastName := "Raj"
	// concetinate 
	fullName := strings.Join([]string{firstName , lastName}," ") // join takes array of string and seperator string 

	fmt.Println("FullName is :" , fullName)
}