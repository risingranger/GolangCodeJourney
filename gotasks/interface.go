package main

import "fmt"

// Define an interface named Shape
type Shape interface {
	Area() float64
}

// Implement the Shape interface for a Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Shape interface for a Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	circ := Circle{Radius: 3}

	shapes := []Shape{rect, circ}

	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
	}
}
