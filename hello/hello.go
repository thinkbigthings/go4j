package main

import (
	"example.com/greetings"
	"fmt"
)

// The main function is the entry point for the application.
// run go run . in the hello directory
func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Jason")
	fmt.Println(message)
}
