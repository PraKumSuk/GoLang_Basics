package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Create file in OS default location i.e. dir is mentioned as ""
	f, err := os.CreateTemp("", "sample") //file name is sample . Note function name is CreateTemp
	check(err)
	fmt.Println("Temp file name:", f.Name())
	defer os.Remove(f.Name())

	//write some bytes to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	//Create dir in OS default location i.e. dir is mentioned as ""
	dname, err := os.MkdirTemp("", "sampledir") //dir name is sampledir. . Note function name is MkdirTemp
	check(err)
	fmt.Println("Temp dir name:", dname)
	defer os.RemoveAll(dname)

	//Join file name with dir we created
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
