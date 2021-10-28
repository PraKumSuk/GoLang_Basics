package main

import "fmt"

func main() {
	jobs := make(chan int, 5) //make buffered channel with 5
	done := make(chan bool)   //make channel

	go func() {
		for {
			j, more := <-jobs //receive msg from jobs channel i.e. loop till you receive msgs
			if more {
				fmt.Println("received job", j) // if job received print received
			} else {
				fmt.Println("received all jobs") // else no jobs... post true msg to done channel
				done <- true
				return
			}
		}
	}()

	//Send 3 jobs to Job channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) //Close Job channel after jobs are complete
	fmt.Println("sent all jobs")

	<-done
}
