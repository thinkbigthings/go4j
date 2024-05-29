package greetings

// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false

const showHints = false

type Shape interface {
	GetArea() float32
	GetPerimeter() float32
}

type Circle struct {
	radius float32
}

func (c Circle) GetArea() float32 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) GetPerimeter() float32 {
	return 2 * 3.14159 * c.radius
}

type Rectangle struct {
	length float32
	width  float32
}

func (r Rectangle) GetArea() float32 {
	return r.length * r.width
}

func (r Rectangle) GetPerimeter() float32 {
	return 2 * (r.length * r.width)
}

// create an appropariate struct to implemnt the Shape interface and build it here
func NewCircle(radius float32) Shape {
	return Circle{
		radius: radius,
	}
}

// create an appropriate struct to implemnt the Shape interface and build it here
func NewRectangle(length, width float32) Shape {
	return Rectangle{
		length: length,
		width:  width,
	}
}

// TODO try composition and see methods be "inherited"

//// This is how your code will be called.
//// Your answer should be the largest value in the numbers array.
//// You can edit this code to try different testing cases.
//circleRadius := float32(5.0)
//rectangleLength := float32(5.0)
//rectangleWidth := float32(6.0)
//shapes := make(map[string]Shape)
//shapes["circle"] = NewCircle(circleRadius)
//shapes["rectangle"] = NewRectangle(rectangleLength, rectangleWidth)
//results := make(map[string]float32)
//if shapes["circle"] != nil {
//results["circleArea"] = shapes["circle"].GetArea()
//results["circlePerimeter"] = shapes["circle"].GetPerimeter()
//} else {
//results["circleArea"] = -1.0
//results["circlePerimeter"] = -1.0
//}
//if shapes["rectangle"] != nil {
//results["rectangleArea"] = shapes["rectangle"].GetArea()
//results["rectanglePerimeter"] = shapes["rectangle"].GetPerimeter()
//} else {
//results["rectangleArea"] = -1.0
//results["rectanglePerimeter"] = -1.0
//}
