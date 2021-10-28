package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n") //write content
	err := os.WriteFile("spk.txt", d1, 0644)
	check(err)

	f, err := os.Create("spk2.txt")
	check(err)
	//close file after opening with defer so u dont miss closing
	defer f.Close()

	//write byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//write string
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	//Flush writes to disk pr stable storage
	f.Sync()

	//buffered writing
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	//flush to write all buffered ops to writer
	w.Flush()

}
