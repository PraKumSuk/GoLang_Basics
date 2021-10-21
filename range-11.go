package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	//range through slice
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range loops through both index and values in slice and array
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range loops through key and value in map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range looping only through key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range looping through Unicode code points of a string
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
