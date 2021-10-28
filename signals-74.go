package main

import (
	"fmt"
	"os"
	"os/signal" //signal package
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1) // Channel to receive OS signal
	done := make(chan bool, 1)      //another channel to notify when to terminate our program

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //registers given channel to notify when signals are received e.g. SIGTERM is a unix signal

	//When signal is received, print and notify program to finish
	//note, this is a blocking go routine
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal") //wait here till the signal is received
	<-done
	fmt.Println("exiting")

	//Note : You cant try this in windows as the signal may not be coming
	//up any time and this program will be waiting forever to get the signal
	//and to exit
}
