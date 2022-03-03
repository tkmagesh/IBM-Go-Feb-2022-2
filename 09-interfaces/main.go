package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

//utilities

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea) {
	fmt.Println("Area =", sa.Area())
}

/* Implement a similar functionality for printing Perimeter */
/*
	perimeter of circle = 2πr
	Perimeter of rectangle = 2 * (height + width)
*/

func main() {
	c := Circle{Radius: 10}
	//fmt.Println("Area =", c.Area())
	PrintArea(c)

	r := Rectangle{Height: 10, Width: 12}
	//fmt.Println("Area =", r.Area())
	PrintArea(r)
}
