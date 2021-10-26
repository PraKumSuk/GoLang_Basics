package main

import "fmt"

func main() {

	// channel buffered to hold 2 values
	messages := make(chan string, 2)

	// adding 2 values to channel
	messages <- "buffered"
	messages <- "channel"

	// no code to receive msgs, but still its avail i.e. buffered channel

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
