package main

import (
	"fmt"
	"sort"
)

//Create a array named byLength  to sort strings by number of characters
type byLength []string

//Implementing sort Interface generic methods our custom way
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"} //input string array
	sort.Sort(byLength(fruits))                   //sort string array byLength
	fmt.Println(fruits)                           //print
}
