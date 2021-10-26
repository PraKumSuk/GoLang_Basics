package main

import "fmt"

func main() {

	//Making a new channel of type string
	messages := make(chan string)

	//  adding msg to i.e. channelName <- msg
	go func() { messages <- "ping" }()

	// retrieving msg from channel i.e. variable := <- channelName
	msg := <-messages
	fmt.Println(msg)
}
