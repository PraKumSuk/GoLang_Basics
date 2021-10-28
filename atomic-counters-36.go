package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64 //unsigned postive integer

	var wg sync.WaitGroup //declare a waiting group

	for i := 0; i < 50; i++ {
		wg.Add(1) //add 1 to wg

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1) //add int with 1, note passing memory address
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
