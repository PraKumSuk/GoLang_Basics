package main

import "fmt"

func main() {

	//Make an empty map of format [key type]value type
	m := make(map[string]int)
	fmt.Println("m is : ", m)

	//Create a map with 2 pairs
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//Assign some values
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	//Get value of key k1
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	//length of map
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//Returns false if there is no key/value present
	//Optional or blank identifer _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

}
