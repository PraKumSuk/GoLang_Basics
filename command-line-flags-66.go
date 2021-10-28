package main

import (
	"flag" //flag package i.e. for command line flags for e.g. flags passed with a cmd wc -l -f -p
	"fmt"
)

func main() {

	//declare a flag named word, default value foo, and a description for this flag
	//Note : flag.String returns a pointer
	wordPtr := flag.String("word", "foo", "a string")

	//declare number and boolean flags too
	//Note : flag.xxxx returns a pointer
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	//Sample showing how to declare a flag,
	//with some other existing variable in our code i.e. in the form of a memory address &
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//Executes the command line parsing
	flag.Parse()

	//Print all the actual values using the pointer i.e. *
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	//To run you can try options like
	//go run .\command-line-flags-66.go -word=opt -numb=7 -fork -svar=flag
	//go run .\command-line-flags-66.go -word=opt a1 a2 a3
}
