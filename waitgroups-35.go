package main

import (
	"fmt"
	"sync"
	"time"
)

//Worker Job
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(2 * time.Second) //wait for 2 seconds say the job is processing
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup // Declare a wait group

	//Start 5 jobs and add each to wait group
	for i := 1; i <= 5; i++ {
		wg.Add(1) //add a job to waitgroup

		i := i

		//wrap worker in a closure, that tells waitgroup that the worker is done
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait() //Block the wait group untill waitgroup counter goes to 0

}
