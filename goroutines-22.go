package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") //sync or seq

	go f("goroutine") //async or concurr

	go func(msg string) {
		fmt.Println(msg)
	}("going") //async

	time.Sleep(time.Second)
	fmt.Println("done")
}
