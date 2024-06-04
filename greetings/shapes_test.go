package greetings

import (
	"testing"
)

//  Ending a file's name with _test.go tells the go test command that this file contains test functions.

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestShape(t *testing.T) {

	// Your answer should be the largest value in the numbers array.
	// You can edit this code to try different testing cases.

	// TODO write a legitimate test case

	circleRadius := float32(5.0)
	rectangleLength := float32(5.0)
	rectangleWidth := float32(6.0)

	shapes := make(map[string]Shape)

	shapes["circle"] = NewCircle(circleRadius)
	shapes["rectangle"] = NewRectangle(rectangleLength, rectangleWidth)

	results := make(map[string]float32)
	if shapes["circle"] != nil {
		results["circleArea"] = shapes["circle"].GetArea()
		results["circlePerimeter"] = shapes["circle"].GetPerimeter()
	} else {
		results["circleArea"] = -1.0
		results["circlePerimeter"] = -1.0
	}
	if shapes["rectangle"] != nil {
		results["rectangleArea"] = shapes["rectangle"].GetArea()
		results["rectanglePerimeter"] = shapes["rectangle"].GetPerimeter()
	} else {
		results["rectangleArea"] = -1.0
		results["rectanglePerimeter"] = -1.0
	}

	testName := t.Name() // this is the test method name
	t.Log(testName)

	//if !want.MatchString(msg) || err != nil {
	//	t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	//}
}

//func TestFindShapes(t *testing.T) {
//	AssertEqual(t, 1, len(FindFruits("purple")))
//	AssertEqual(t, 2, len(FindFruits("red")))
//	AssertEqual(t, 3, len(FindFruits("yellow")))
//}

//func TestFilterShapes(t *testing.T) {
//	AssertEqual(t, 2, len(FilterFruits()))
//}
