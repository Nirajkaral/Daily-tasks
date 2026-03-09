package main

import (
	"fmt"
	"math"
)

// Interface
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

// Rectangle implementation
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle implementation
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func poly() {

	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 7}

	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Circle Area:", c.Area())
}