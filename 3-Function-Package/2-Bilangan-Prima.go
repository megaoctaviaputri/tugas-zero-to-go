package main

import "fmt"

func Prima(N int) string {

	var count int
	var hasil string

	for i := 1; i <= N; i++ {
		if N%i == 0 {
			count++
		}
	}

	if count == 2 {
		hasil = "Bilangan Prima"
	} else {
		hasil = "Bukan Bilangan Prima"
	}
	
	return hasil
}

func main() {
	fmt.Println(Prima(1))
	fmt.Println(Prima(3))
	fmt.Println(Prima(7))
	fmt.Println(Prima(6))
	fmt.Println(Prima(23))
}
