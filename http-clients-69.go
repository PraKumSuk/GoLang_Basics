package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	//Create http.Client object and call its Get Method i.e.
	//http.Get()
	resp, err := http.Get("http://www.go.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//Print response
	fmt.Println("Response status:", resp.Status)

	//Print  only first 5 lines of response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
