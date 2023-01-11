package main

import "fmt"

type square struct {
	height float64
	width  float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{
		height: 10,
		width:  10,
	}
	t := triangle{
		base:   15,
		height: 8,
	}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.height * s.width
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
