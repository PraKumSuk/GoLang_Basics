package main

import "fmt"

func main() {

	//simplest for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//traditional for loop without braces
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//infinite loop with break
	for {
		fmt.Println("loop")
		break
	}

	//loop with continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
