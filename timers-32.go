package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(5 * time.Second) //Set timer for 5 seconds

	//Print msg when timer1 is fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	//Wait for 2 seconds and stop the timer, before it starts firing at 5th second
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(5 * time.Second)
}
