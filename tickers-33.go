package main

import (
	"fmt"
	"time"
)

func main() {

	//ticker keeps firing untill stopped
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	//Print msg everytime there is a ticker event
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	//Wait for 10 seconds and then stop the ticker
	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- true

	//print ticket stopped
	fmt.Println("Ticker stopped")
}
