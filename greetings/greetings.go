package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// A function whose name starts with a capital letter can be called by a function not in the same package.
// This is known in Go as an exported name.

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	// Go's arrays are values.
	// An array variable denotes the entire array; it is not a pointer to the first array element.
	// This means that when you assign or pass around an array value you will make a copy of its contents.
	// (To avoid the copy you could pass a pointer to the array, but then that’s a pointer to an array, not an array.)
	// One way to think about arrays is as a sort of struct but with indexed rather than named fields:
	// a fixed-size composite value.

	// An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types)

	// Arrays have their place, but they are a bit inflexible, so you don't see them too often in Go code.
	// Slices, though, are everywhere.

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("name parameter is empty but it should be a non-empty string")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package
//(in other words, it's not exported).

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {

	// A slice of message formats.
	// When declaring a slice, you omit its size in the brackets, like this: []string.
	//This tells Go that the size of the array underlying the slice can be dynamically changed.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
