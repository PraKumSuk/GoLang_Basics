package main

import (
	"fmt"
	"io"
	"os/exec" //exec package to exec or control other OS processes
)

func main() {

	//exec.Command creates an object that represents external process and picks date from Sys
	//dateCmd := exec.Command("date") //Does not work in windows, so use below
	dateCmd := exec.Command("cmd.exe") //If above works comment this line

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	//Another example to perform a grep,
	//pass some input using stdin
	//wait till the grep command is processed
	//stdout the response received
	//and then wait for process to exit
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	//spawning cmd passed along with our program that
	//appears as if its just passed in one command line string
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
