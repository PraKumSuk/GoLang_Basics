package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//buffered io to scan keyboard input
	scanner := bufio.NewScanner(os.Stdin)

	//scan for next lines and convert to uppercase
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	//check for err if any
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
