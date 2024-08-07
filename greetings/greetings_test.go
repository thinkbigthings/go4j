package greetings

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

//  Ending a file's name with _test.go tells the go test command that this file contains test functions.

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {

	testName := t.Name() // this is the test method name
	t.Log(testName)
	//log.Println(testName)
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestFindFruits(t *testing.T) {
	assert.Equal(t, 1, len(FindFruits("purple")))
	assert.Equal(t, 2, len(FindFruits("red")))
	assert.Equal(t, 3, len(FindFruits("yellow")))
}

func TestFilterFruits(t *testing.T) {
	assert.Equal(t, 2, len(FilterFruits()))
}
