package main

import (
	"crypto/sha1" //other hash algorithms like md5 can also be used
	"fmt"
)

func main() {

	s := "sha1 this string"

	h := sha1.New()    //to generate a sha1
	h.Write([]byte(s)) //convert str to bytes i.e. []bytes(s), and then get sha1 byte
	bs := h.Sum(nil)   //getting final hash as a byte slice

	//Print sha in human readable hex form normally
	fmt.Println(s)

	fmt.Printf("%x\n", bs)
}
