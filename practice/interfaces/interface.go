package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func newRectangle(width, height float64) *rect {
	r := rect{
		height: height,
		width:  width,
	}
	return &r
}

type circle struct {
	radius float64
}

func newCircle(radius float64) *circle {
	c := circle{
		radius: radius,
	}
	return &c
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := newRectangle(10, 10)
	c := newCircle(20)

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

}
