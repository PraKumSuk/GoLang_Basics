package main //Test code shud normally reside in same package it tests code for

import (
	"fmt"
	"testing" //package testing
)

//func to give smallest number
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//test to test above func
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans) //throw error
	}
}

//test multiple scenarios in a single test
func TestIntMinTableDriven(t *testing.T) {

	//Set of input values and expected output
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	//looping through to determine test results
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
