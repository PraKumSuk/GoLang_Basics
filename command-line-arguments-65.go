package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args        //returns command line arguments passed
	argsWithoutProg := os.Args[1:] //refers 1st value in slice

	arg := os.Args[3] //number of arguments to be fetched

	fmt.Println(argsWithProg)    //prints 1st arg i.e. path_of_program
	fmt.Println(argsWithoutProg) //prints args, apart from path_of_program
	fmt.Println(arg)

	//Note : to run this give more than 3 arguments so it does not fail for e.g
	//go run .\command-line-arguments-65.go a c d s d

}
