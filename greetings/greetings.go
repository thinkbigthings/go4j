package greetings

import "fmt"

// A function whose name starts with a capital letter can be called by a function not in the same package.
// This is known in Go as an exported name.

// Hello returns a greeting for the named person.
func Hello(name string) string {

	// The := operator is a shortcut for declaring and initializing a variable in one line
	// Go uses the value on the right to determine the variable's type
	// So there is type inference, just like in Java

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
