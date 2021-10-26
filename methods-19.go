package main

import "fmt"

type rect struct {
	width, height int
}

//method format is
//func (inputParams) methodName() returnType{}
func (r *rect) area() int { //Input param variable name is r of struct type react, and is a pointer
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
