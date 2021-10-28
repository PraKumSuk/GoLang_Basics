package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println("Unix time is ", now) //Prints the epoch time i.e. 2021-10-27 13:20:21.2596113 +0530 IST m=+0.003429301

	//Since no milli second avail in epoch, we divide nanoseconds with 1000000
	millis := nanos / 1000000
	fmt.Println(secs)   //1635321021
	fmt.Println(millis) //1635321021259
	fmt.Println(nanos)  //1635321021259611300

	//convert integer seconds to nanoseconds
	fmt.Println(time.Unix(secs, 0))  //2021-10-27 13:20:21 +0530 IST
	fmt.Println(time.Unix(0, nanos)) //2021-10-27 13:20:21.2596113 +0530 IST
}
