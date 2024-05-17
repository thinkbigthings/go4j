package main

// Go programs are constructed by linking together packages.
// A package in turn is constructed from one or more source files that together declare constants, types, variables and
// functions belonging to the package and which are accessible in all files of the same package.
// Those elements may be exported and used in another package.

// package imports
import (
	"example.com/greetings"
	"fmt"
	"log"
)

// The main function is the entry point for the application.
// run go run . in the hello directory
func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Jason")
	exitFatallyOn(err)
	fmt.Println(message)

	names := []string{"Jason", "Wendy", "Jacob", "Darin"}
	messages, err := greetings.Hellos(names)
	exitFatallyOn(err)
	fmt.Println(messages)

	// Get an error message and print it.
	message, err = greetings.Hello("")
	exitFatallyOn(err)
	fmt.Println(message)
}

func exitFatallyOn(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
