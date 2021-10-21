package main

import "fmt"

func main() {

	var a = "initial" //Declare and Initialised
	fmt.Println(a)

	var b, c int = 1, 2 //declared and initialise multiple variables i.e. format is var varName valueType = actualValue
	fmt.Println(b, c)   //prints b and c values i.e. 1 2

	var d = true
	fmt.Println(d)

	var e int      //Zero Valued i.e. Declared but not initiliased
	fmt.Println(e) //Prints default value i.e. 0

	f := "apple" //var f string = "apple" can be written directly as f := "apple"
	fmt.Println(f)
}
