package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area()
}

type Circle struct {
	name   string
	radius float64
}

type Square struct {
	name string
	side int
}

func (c Circle) area() {
	fmt.Println(math.Pi * c.radius * c.radius)
}

func (s Square) area() {
	fmt.Println(s.side * s.side)
}

func printArea(shape Shape) {
	shape.area()
}

func main() {
	circle := Circle{name: "Ankit", radius: 10}
	square := Square{name: "Sir", side: 5}

	printArea(circle)
	printArea(square)
}
