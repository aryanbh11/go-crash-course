package main

import (
	"fmt"
	"math"
)

// Define Interface 
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	b, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) area() float64 {
	return 0.5 * t.b * t.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{radius:1.2}
	triangle := Triangle{b:1.5, h:1.6}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(triangle))
}