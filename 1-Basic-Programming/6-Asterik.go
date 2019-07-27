package main

import "fmt"

func main(){
	var input = 5
	fmt.Println(input)
	for i := 0; i < input; i++ {
		for j := 0; j < input-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i ; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}