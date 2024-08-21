package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 10}
	rectangle := Rectangle{length: 5, width: 8}

	fmt.Printf("Area of the circle is: %.2f\n", getArea(circle))
	fmt.Printf("Area of the rectangle is: %.2f\n", getArea(rectangle))
}
