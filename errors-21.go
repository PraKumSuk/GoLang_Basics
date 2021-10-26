package main

import (
	"errors"
	"fmt"
)

//func f1, input param arg, returs int, error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42") //new error //if arg is 42 throw error
	}
	return arg + 3, nil //nil means no error occurred //else add 3 and return
}

//argError of type struct
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} //if arg is 42 throw error
	}
	return arg + 3, nil //else add 3 and return
}

func main() {

	for _, i := range []int{7, 42} { // loop through the array i.e. for values 7,42
		if r, e := f1(i); e != nil { // r, e are return types. r holds value e holds error.
			fmt.Println("f1 failed:", e)
		} else { //if no error print worked
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e) //prints error thrown here
		} else {
			fmt.Println("f2 worked:", r) //prints worked
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
