package main

import (
	"fmt"
	"time"
)

func main() {

	//channel of 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//Setting limit of 5 seconds
	//Range through requests, wait for the limiter
	limiter := time.Tick(2 * time.Second)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//Setting bursty limiter to burst process 3 requests at a time i.e. concurrently
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(10 * time.Second) {
			burstyLimiter <- t
		}
	}()

	//Setting burstyRequests of 5, so 1-3 happens concurrently and 4th and 5th again happens after 10 sec
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	//Print bursty requests
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
