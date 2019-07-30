package main

import "fmt"

func main(){
	input := []int{10, 4, 5, 2, 4}
	frekuensiBilangan := map[int]int{}
	var nilaiMaks int
	var modus int

	for _, bilangan := range input {
		frekuensiBilangan[bilangan]++
		if frekuensiBilangan[bilangan] > nilaiMaks {
			nilaiMaks = frekuensiBilangan[bilangan]
			modus = bilangan
		}
	}

	if nilaiMaks == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(modus)
	}
}