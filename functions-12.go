package main

import "fmt"

//Simple function of format func funcName(var type, var type) returnType{}
func plus(a int, b int) int {
	return a + b
}

//for multiple similar params mention type only once
func plusPlus(a, b, c int) int {
	return a + b + c //return type is mandatory
}

//main function runtime.main_main case sensitive cannot be changed
func main() {

	res := plus(1, 2) //Invoking a function normally
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
