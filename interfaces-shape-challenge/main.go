package main

import "fmt"

type Square struct {
	sideLength float64
}
type Triangle struct {
	height float64
	base   float64
}

type Shape interface {
	getArea() float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64 {
	return t.height * t.base / 2
}

func PrintArea(sh Shape) {
	fmt.Println("Area is: ", sh.getArea())
}

func main() {
	s := Square{
		sideLength: 10,
	}
	PrintArea(s)
	t := Triangle{
		height: 10,
		base:   10,
	}
	PrintArea(t)
}
