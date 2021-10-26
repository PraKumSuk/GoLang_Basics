package main

import "fmt"

//struct defined
type person struct {
	name string
	age  int
}

//func accepting pointer and updates age
//and returns same
func newPerson(name string) *person {

	p := person{name: name} //new struc is being created here
	p.age = 42
	return &p //returning local p struct object so original passed will be intact
}

func main() {

	fmt.Println(person{"Bob", 20}) //new struct instantiation

	fmt.Println(person{name: "Alice", age: 30}) //new struct with name and age

	fmt.Println(person{name: "Fred"}) //new struct with name and zero valued default to age

	fmt.Println(&person{name: "Ann", age: 40}) //original value in tact though pointer is passed

	fmt.Println(newPerson("Jon")) //actual object itself is passed, so adds age 42

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s //making a copy of s struct i.e. pass by memory address
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
