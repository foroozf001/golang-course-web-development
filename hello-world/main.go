package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page\n")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page, sum:%d\n", sum))
}

// About is the about page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32 = 100.0, 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided %f by is %f\n", x, y, f))
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// divideValues adds two integers and returns the sum
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		log.Println(err)
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
