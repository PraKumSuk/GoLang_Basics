package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Panic if any error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//read
	dat, err := os.ReadFile("spk.txt")
	check(err)
	fmt.Print(string(dat))

	//open
	f, err := os.Open("spk.txt")
	check(err)

	//read 5 bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	//read from 6
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	//read atleast
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	//buffered reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	//close file
	f.Close()
}
