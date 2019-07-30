package main

import "fmt"

func main(){
	input := []int{10, 4, 5, 2, 4}
	frekuensiBilangan := map[int]int{}
	var mayor int

	for _, bilangan := range input {
		frekuensiBilangan[bilangan]++
		if frekuensiBilangan[bilangan] > (len(input) / 2) {
			mayor = bilangan
		}
	}

	if mayor > 0 {
		fmt.Println(mayor)
	} else {
		fmt.Println("Tidak Ditemukan")
	}
}