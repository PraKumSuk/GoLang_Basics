package main

import "fmt"

func main() {

	//slice no size needs to be mentioned
	a := []int{1, 2, 3}
	fmt.Println("first slice is : ", a)

	//making slice using make method
	s := make([]string, 3)
	fmt.Println("emp:", s) //empty slice zero value is NIL or Blank

	//Setting 3 values to slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    //Prints [a b c]
	fmt.Println("get:", s[2]) //Prints c i.e. value at index of 2 Note index starting from 0 as array

	fmt.Println("len:", len(s)) //Prints 3 i.e. length of slice is 3

	s = append(s, "d")      //Append d to slice
	s = append(s, "e", "f") //Append multiple values to slice
	fmt.Println("apd:", s)  //Prints [a b c d e f]

	c := make([]string, len(s)) //Making new slice same as s
	copy(c, s)                  //Copy s to c
	fmt.Println("cpy:", c)      //Print c slice i.e. prints same as s i.e. [a b c d e f]

	//Slicing further
	l := s[2:5]            //index 2 to 4 i.e. when 5th dont add
	fmt.Println("sl1:", l) //Prints c,d,e

	l = s[:5] //Default 0, so print 0 to 4
	fmt.Println("sl2:", l)

	l = s[2:] //Default length, so print index 2 to last or length of slice
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//2 dimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
