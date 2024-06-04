package greetings

import (
	"math"
)

// Point2D methods on both value and pointer receivers is not recommended by the Go Documentation.
// let's use only pointer receivers for the methods.
type Point2D struct {
	X float32
	Y float32
}

func (p *Point2D) GetCenter() Point2D {
	return *p
}

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

func (c *Circle) GetArea() float32 {
	return 3.14159 * c.radius * c.radius
}

func (c *Circle) GetPerimeter() float32 {
	return 2 * 3.14159 * c.radius
}

type Rectangle struct {
	Point2D
	length float32
	width  float32
}

func (r *Rectangle) GetArea() float32 {
	return r.length * r.width
}

func (r *Rectangle) GetPerimeter() float32 {
	return 2 * (r.length * r.width)
}

func NewCircle(radius float32) Shape {
	return &Circle{
		radius: radius,
	}
}

func NewRectangle(length, width float32) Shape {
	return &Rectangle{
		length: length,
		width:  width,
	}
}
