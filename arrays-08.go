package main

import "fmt"

func main() {

	var a [5]int           //array of size 5 declared but not initialised
	fmt.Println("emp:", a) //Prints whole array, and all 5 elements will have 0 by default

	a[4] = 100 //setting value 100 to a[4]
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) //len(arrayName) prints length of array

	b := [5]int{1, 2, 3, 4, 5} //declare and initialise array with values
	fmt.Println("dcl:", b)

	// 2 dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
