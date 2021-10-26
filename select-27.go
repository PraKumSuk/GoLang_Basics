package main

import (
	"fmt"
	"time"
)

func main() {

	//creating 2 channels
	c1 := make(chan string)
	c2 := make(chan string)

	//channel 1 to add a msg after 1 sec
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	//channel 2 to add a msg after 2 sec
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		//Select selects each case as and when it receives a msg
		//from channel
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
