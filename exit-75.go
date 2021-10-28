package main

import (
	"fmt"
	"os"
)

func main() {

	//Defer will not be executed when a os.Exit is called,
	//so os.Exit has the highest priority
	defer fmt.Println("!")

	os.Exit(3) //Immediately exit, with status 3
}
