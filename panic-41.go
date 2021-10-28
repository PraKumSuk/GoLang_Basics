package main

import "os"

func main() {

	// panic("a problem") //If this panic is enabled, below becomes unreachable

	_, err := os.Create("spkFile.txt") //creates file in the current dir
	if err != nil {
		panic(err)
	}
}
