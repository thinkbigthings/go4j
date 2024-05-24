package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// A function whose name starts with a capital letter can be called by a function not in the same package.
// This is known in Go as an exported name.

// FindFruits returns a list of fruits of the desired color.
func FindFruits(desiredColor string) []string {

	// https://github.com/jucardi/go-streams

	// Create a map with string keys and string values
	// this is a map literal
	fruitColor := map[string]string{
		"peach":      "orange",
		"apple":      "red",
		"pear":       "yellow",
		"plum":       "purple",
		"pineapple":  "yellow",
		"banana":     "yellow",
		"strawberry": "red",
		"orange":     "orange",
	}

	// range over a map iterates over key/value pairs, there is no index like range over a slice
	found := make([]string, 0)
	for key, val := range fruitColor {
		if val == desiredColor {
			found = append(found, key)
		}
	}

	return found

	// ConditionalFunc is an alias to `func(interface{}) bool`
	// which serves to define if a condition is met for an element in the collection.
	//type ConditionalFunc func(interface{}) bool
	//isDesiredColor := func(e interface{}) bool {
	//	tuple, ok := e.(types.Tuple)
	//	if !ok {
	//		return false
	//	}
	//	return tuple.At(1).String() == desiredColor
	//}

	// Filter the map to keep only the desired fruits
	//desiredKeys := streams.FromMap(fruitColor).Filter(isDesiredColor).ToArray()

}

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

// Hellos returns a map that associates each of the named people
// with a greeting message.

func Hellos(names []string) (map[string]string, error) {

	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	// the underscore (_) is known as the "blank identifier." It is used to discard values that you don't need
	// in this case, index.
	// This loop is like Java's enhanced for loop but the index is made available, and we can ignore it with _
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package
//(in other words, it's not exported).

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {

	// A slice of message formats.
	// When declaring a slice, you omit its size in the brackets, like this: []string.
	// This tells Go that the size of the array underlying the slice can be dynamically changed.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	dynamicFormats := make([]string, 0)

	for _, f := range formats {
		dynamicFormats = append(dynamicFormats, f)
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return dynamicFormats[rand.Intn(len(dynamicFormats))]
}
