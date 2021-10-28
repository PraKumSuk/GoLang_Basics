package main

import (
	"fmt"
	"sort" //sort module
)

func main() {

	//Sort string array
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	//Sort int array
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	//Check whether sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
