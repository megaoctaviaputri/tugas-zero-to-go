package main

import "fmt"

func main() {

	// var sliceInput = []float64{1, 2, 3, 4, 5} // test case 1
	// var sliceInput = []float64{1, 3, 4, 10, 12, 13} // test case 2
	var sliceInput = []float64{7, 7, 8, 8} // test case 3
	var median float64
	var temp int

	if len(sliceInput)%2 == 0 {
		temp = len(sliceInput) / 2
		median = (sliceInput[temp-1] + sliceInput[temp]) / 2
	} else {
		temp = len(sliceInput) / 2
		median = sliceInput[temp]
	}
	fmt.Println(median)
}