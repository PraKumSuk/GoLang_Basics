package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()

	p(now) //current time

	//Some past time, time is always associated with timezone
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year()) //prints year
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond()) //prints nano seconds i.e. 651387237
	p(then.Location())   //prints location i.e. time zone i.e. UTC
	p(then.Weekday())    //prints day i.e. Tuesday

	//Compare time
	p(then.Before(now)) //is then before now, prints true i.e. past time
	p(then.After(now))  //false
	p(then.Equal(now))  //false

	//Diff between 2 times
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())   //diff in hours
	p(diff.Minutes()) //diff in minutes
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//Add/Sub time i.e. advance/go back in time by certain duration
	p(then.Add(diff))  //Add some more time to given time
	p(then.Add(-diff)) //Subtract some more time to given time
}
