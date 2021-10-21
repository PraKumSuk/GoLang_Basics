package main

import "fmt"

//func with 2 return types
func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals() //receiving 2 return values
	fmt.Println(a)
	fmt.Println(b)

	//func returns 2 values but you are interested only in 1 value,
	//which ever is not required can be replaced with
	//blank identifier _, c means 1st return type vl be blanked out
	_, c := vals()
	fmt.Println(c)
}
