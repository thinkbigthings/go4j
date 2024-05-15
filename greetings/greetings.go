package greetings

import (
	"errors"
	"fmt"
)

// A function whose name starts with a capital letter can be called by a function not in the same package.
// This is known in Go as an exported name.

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("name parameter is empty but it should be a non-empty string")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
