package main

import (
	"fmt"
	"math"
)

func main() {

	// var array = [5]int{1, 2, 3, 4, 5} // test case 1
	// var array = [5]int{3, 5, 7, 5, 3} // test case 2
	var array = [5]int{6, 5, 4, 7, 3} // test case 3
	var sum int
	var mean float64

	for i := 0; i < len(array); i++ {
		sum = sum + array[i]
	}

	mean = sum / float64(len(array))
	fmt.Println(math.Round(mean))

}