package main

import (
	"fmt"
	"math"
)

//interface created with method signatures
type geometry interface {
	area() float64
	perim() float64
}

//creating 2 struc for rect and circle
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//interface function implementation
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//interface function implementation
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//a typical method measure
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
