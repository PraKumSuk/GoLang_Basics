package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	//recover will always be called from within defer
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic() //Enabling this will cause the program to not run

	fmt.Println("After mayPanic()")
}
