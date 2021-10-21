package main

import "fmt"

//func with multiple parameters i.e. similar to varargs in java
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)    //passing 2 params to sum
	sum(1, 2, 3) //passing 3 parsms to sum Another example of variadic function is fmt itself

	nums := []int{1, 2, 3, 4} //assigning slice of values to a variadic function as input param
	sum(nums...)
}
