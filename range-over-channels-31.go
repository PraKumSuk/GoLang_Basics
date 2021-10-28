package main

import "fmt"

func main() {

	queue := make(chan string, 3) //create a buffered channel of size 2
	queue <- "one"                //add a msg
	queue <- "two"                //add another msg
	close(queue)                  // close the channel
	// queue <- "threee" 			  //not permitted, u can only read after closing but not write

	//You can still read msgs from channel even after closing it. i.e.
	//you can still receive msgs after closing the channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
