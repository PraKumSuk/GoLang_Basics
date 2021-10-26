package main

import "fmt"

//a recursive function calling itself
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) //calling itself here
}

func main() {
	fmt.Println(fact(7)) //Prints factorial of 7 i.e. 7 * 6 * ...

	//fibonacci function calling itself multiple times
	//also func defined and declared to a variable
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)

	}

	fmt.Println(fib(7))
}
