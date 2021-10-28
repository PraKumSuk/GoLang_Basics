package main

import (
	"fmt"
	"net/http"
	"time"
)

//This prints or logs every request once received, processed and completed
func hello(w http.ResponseWriter, req *http.Request) {

	/* A context.Context is created for each request by the net/http machinery,
	and is available with the Context() method. */
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	//wait for 10seconds for server to start
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		//if closing, print to console the reason of closing the reply to client
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)

	//Start the program and try http://localhost:8090/hello
	//and then check the console logs

}
