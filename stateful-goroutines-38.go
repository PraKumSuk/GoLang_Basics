package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// struct to hold read op details
type readOp struct {
	key  int      //key to read some value of a map
	resp chan int //post read value as response in channel
}

// struct to hold write op details
type writeOp struct {
	key  int       //key to write some value
	val  int       //value to be written
	resp chan bool //post value stating write comleted
}

func main() {

	var readOps uint64  //count readOps
	var writeOps uint64 //count writeOps

	reads := make(chan readOp)   //readOp channel
	writes := make(chan writeOp) //writeOp channel

	//func to either perform read/write opn accordingly
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key] //read value and post in read channel
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true //write value and post in write channel
			}
		}
	}()

	//go routine to perfor some  writes
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//go routine to perform some reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second) //wait fora Second to complete go routines

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
