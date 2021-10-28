package main

import (
	"fmt"
	"net/http" //package
)

//Declare a hello Hander i.e. a Handler is an object implementing http.Handler interface
//func serving as handlers take http.ResponseWriter and htt.Request as arguments
//A Hello Handler
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

//Read all headers
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	//Register our hello handler on default server routes using
	//http.HandleFunc
	http.HandleFunc("/hello", hello)     //Path to access or endpoint
	http.HandleFunc("/headers", headers) //Path to access or endpoint

	//ListenAndServe serves request, with default router handler.nil
	http.ListenAndServe(":8090", nil)

	//Try running the program and hitting http://localhost:8090/hello or http://localhost:8090/headers
}
