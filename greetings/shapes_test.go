package greetings

import (
	"math"
	"testing"
)

//  Ending a file's name with _test.go tells the go test command that this file contains test functions.

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestShape(t *testing.T) {

	circleRadius := float32(5.0)
	rectangleLength := float32(5.0)
	rectangleWidth := float32(6.0)

	// create a map of shapes to demonstrate polymorphic behavior
	shapes := make(map[string]Shape)

	shapes["circle"] = NewCircle(circleRadius)
	shapes["rectangle"] = NewRectangle(rectangleLength, rectangleWidth)

	results := make(map[string]float32)
	results["circleArea"] = shapes["circle"].GetArea()
	results["circlePerimeter"] = shapes["circle"].GetPerimeter()
	results["rectangleArea"] = shapes["rectangle"].GetArea()
	results["rectanglePerimeter"] = shapes["rectangle"].GetPerimeter()

	AssertEqual(t, float32(22.0), results["rectanglePerimeter"])
	AssertTrue(t, math.Abs(10.0*math.Pi-float64(results["circlePerimeter"])) < 0.0001)
}
