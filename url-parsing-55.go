package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	//sample url
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	//if invalid char panic
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) //print scheme

	//User variable contains both un and pwd
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	//Host name contains both host and port
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	//Print path i.e. actual path after host and port
	fmt.Println(u.Path)

	//Print fragment i.e. anything after #
	fmt.Println(u.Fragment)

	//Get all query params to put in map if required
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)

	fmt.Println(m["k"][0]) //print 1st query param or key value pair
}
