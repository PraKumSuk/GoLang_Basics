package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) { //channel here of type boolean
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true //add true to channel
}

func main() {

	done := make(chan bool, 1)
	go worker(done) //this is a go routine, works concurrently

	// this line holds program untill above line completes, if this line is not present.
	//Program would complete even before above line completes as its asyn
	<-done //refernce to existing channel, waits untill there is a msg

}
