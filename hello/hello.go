package main

import (
	"example.com/greetings"
	"fmt"
)

// The main function is the entry point for the application.
// run go run . in the hello directory
func main() {

	// Get a greeting message and print it.
	message, err := greetings.Hello("Jason")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Exiting")
		return
	}
	fmt.Println(message)

	// Get an error message and print it.
	message, err = greetings.Hello("")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Exiting")
		return
	}
	fmt.Println(message)
}
