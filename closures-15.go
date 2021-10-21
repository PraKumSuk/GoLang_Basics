package main

import "fmt"

//Closures i.e. a sub function or inline function is returned
//which can be put in a variable
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//State is Persisted
	nextInt := intSeq()    //Calling first time, this returns a function itself, which is stored in nextInt
	fmt.Println(nextInt()) //Invoking returned function with its variable along with (), prints 1
	fmt.Println(nextInt()) //.....prints 2
	fmt.Println(nextInt()) //.......prints 3

	//New state
	newInts := intSeq()    //Calling again for a new variable i.e. first time, returns new function again
	fmt.Println(newInts()) //prints 1
}
