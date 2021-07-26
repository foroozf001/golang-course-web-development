package main

import (
	"fmt"
	"log"
)

func main() {
	// section_VariablesAndFunctions()
	// section_Pointers()
}

func section_VariablesAndFunctions() {
	fmt.Println("Hello, world!")
	var whatToSay string
	var i int
	whatToSay = "Goodbye, cruel world."
	fmt.Println(whatToSay)
	i = 7
	fmt.Println("i value is", i)
	whatWasSaid := saySomething()
	fmt.Println("The function returnd", whatWasSaid)
}

func saySomething() string {
	return fmt.Sprint("Something")
}

func section_Pointers() {
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
