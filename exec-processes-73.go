package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("ls") //ls command is given here
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"} //pass ls command with other flags here

	env := os.Environ() //getting environment variables

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	//This will not work in windows as these are unix/linux based cmds
}
