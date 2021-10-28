package main

import (
	"fmt"
	"time"
)

// Worker Job
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(2 * time.Second) //Sleep for 2 seconds, say a job could take 2 seconds for example
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5                  // Number of Jobs
	jobs := make(chan int, numJobs)    //Job
	results := make(chan int, numJobs) // Job result

	//Starting 3 Jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//Receiving job results of 5 Jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) //Close Job

	//Wait untill all jobs have completed
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
