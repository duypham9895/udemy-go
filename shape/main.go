package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	s := square{sideLength: 3}
	t := triangle{height: 3, base: 5}

	printArea(s)
	printArea(t)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(sp shape) {
	fmt.Println("Area of this shape is ", sp.getArea())
}
