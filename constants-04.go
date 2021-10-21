package main

//Importing multiple modules
import (
	"fmt"
	"math"
)

const s string = "constant" //a string  constant

//Function starts here

func main() {
	fmt.Println(s)

	const n = 500000000 //type is not mandatory and const can be declared anywhere

	const d = 3e20 / n //Allows expressions
	fmt.Println(d)

	fmt.Println(int64(d)) //converts to int64 as int64 expects int

	fmt.Println(math.Sin(n)) //converts to float, as math.Sin expects float
}
