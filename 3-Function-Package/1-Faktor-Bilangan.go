package main

import "fmt"

func faktorBilangan(N int) {

	for i := 1; i <= N; i++ {
		if N%i == 0 {
			fmt.Print(i, " ")
		}
	}
	
	fmt.Println()
}

func main() {
	faktorBilangan(6)
	faktorBilangan(12)
	faktorBilangan(14)
	faktorBilangan(16)
	faktorBilangan(20)
}
