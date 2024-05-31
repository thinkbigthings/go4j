package greetings

import (
	"math"
)

type Point2D struct {
	X float32
	Y float32
}

func (p Point2D) GetCenter() Point2D {
	return p
}

// TODO make consistent usage and test it out
//
//	methods on both value and pointer receivers is not recommended by the Go Documentation.
func (p *Point2D) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point2D) Scale(f float32) {
	p.X = p.X * f
	p.Y = p.Y * f
}

type Position interface {
	GetCenter() Point2D
}

type Shape interface {
	Position
	GetArea() float32
	GetPerimeter() float32
}

type Circle struct {
	Point2D
	radius float32
}

func (c Circle) GetArea() float32 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) GetPerimeter() float32 {
	return 2 * 3.14159 * c.radius
}

type Rectangle struct {
	Point2D
	length float32
	width  float32
}

func (r Rectangle) GetArea() float32 {
	return r.length * r.width
}

func (r Rectangle) GetPerimeter() float32 {
	return 2 * (r.length * r.width)
}

func NewCircle(radius float32) Shape {
	return Circle{
		radius: radius,
	}
}

func NewRectangle(length, width float32) Shape {
	return Rectangle{
		length: length,
		width:  width,
	}
}

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
