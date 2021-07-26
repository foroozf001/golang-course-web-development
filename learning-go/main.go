package main

import "fmt"

func main() {
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
