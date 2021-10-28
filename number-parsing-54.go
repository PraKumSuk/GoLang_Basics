package main

import (
	"fmt"
	"strconv"
)

//Convert string to numbers
func main() {

	//convert string to float with 64bit precision
	f, _ := strconv.ParseFloat("1.234", 64) //note blank operator to hold error if any
	fmt.Println(f)

	//convert string to int
	i, _ := strconv.ParseInt("123", 0, 64) //0 means infer base from string, 64 bit means result shud be of 64bit size
	fmt.Println(i)

	//convert string to hex format number
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//base 10
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	//base 10 throwing an error
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
