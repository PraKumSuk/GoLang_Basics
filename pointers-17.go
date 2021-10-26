package main

import "fmt"

//Copy by value here
//i.e. a new copy of value will be assigned
//when this function is invoked
func zeroval(ival int) {
	ival = 0
}

//Assigning a pointer variable of type int
// i.e. *int
func zeroptr(iptr *int) { //pass pointer or memory address
	*iptr = 0 //update actual value for the given pointer/memory address
}

func main() {
	i := 1
	fmt.Println("initial:", i) //prints 1

	zeroval(i)
	fmt.Println("zeroval:", i) //prints 1 again, since zeroval(i) is pass by value or copy is passed to the method

	// to get memory address of a pointer variable
	// use &pointerVariableName
	zeroptr(&i)                //pass by pointer or memory address, so this pointer value is being set as 0 in zeroptr function
	fmt.Println("zeroptr:", i) //so it now prints 0, as the initial value itself was changed

	fmt.Println("pointer:", &i) //rints pointer memory address 0xc000012088
}
