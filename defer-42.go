package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("defer.txt") //Creating a file
	defer closeFile(f)           //defer invoking closeFile function i.e. do this finally
	writeFile(f)                 //write file
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	//Capture if ther is any eror while trying to close the file
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
