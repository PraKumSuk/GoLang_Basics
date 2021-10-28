package main

import (
	"fmt"
	"strings"
)

//returns int index of given substring
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//returns bool if given substring is present
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// returns true if any one of the items contain given subset
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// returns true only if all items contains the given subset
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//Filter all items that contain the subset given or matching
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//Map returns a new slice after applying a function or modifying the data
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear")) //prints 2

	fmt.Println(Include(strs, "grape")) //returns false

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p") //returns true
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p") //returns false
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}
