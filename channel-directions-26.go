package main

import "fmt"

//channel that only accepts msg
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//Pong function that posts msg to pongs channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message") //channel that accepts msg
	pong(pings, pongs)            //channel that accepts msg
	fmt.Println(<-pongs)          //channel that returns msg
}
