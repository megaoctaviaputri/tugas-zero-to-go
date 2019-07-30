package main

import "fmt"

func DeretGeometri(arr []int) bool {

	var pengali int = arr[1] / arr[0]
	var count int
	
	for i := 1; i < len(arr); i++ {
		if arr[i]/pengali == arr[i-1] {
			count++
		}
	}

	if count == len(arr)-1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(DeretGeometri([]int{1, 3, 9, 27, 81}))
	fmt.Println(DeretGeometri([]int{2, 4, 8, 16, 32}))
	fmt.Println(DeretGeometri([]int{2, 4, 6, 8}))
	fmt.Println(DeretGeometri([]int{2, 6, 18, 54}))
	fmt.Println(DeretGeometri([]int{1, 2, 3, 4, 7, 9}))
}
