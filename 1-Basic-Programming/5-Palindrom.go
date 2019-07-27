package main

import (
	"fmt"
	"strings"
)

func main(){
	var input = "kasur rusak"
	var isPalindrome = 1
	inour = strings.ToLower(input)
	for i := 0; i < len(input); i++ {
		if input[i] == input[len(input)-1-i] {
			isPalindrome *= 1
		} else {
			isPalindrome *= 0
		}
	} 
	if isPalindrom == 1 {
		fmt.Println("Output: Palindrom")
	} else {
		fmt.Println("Output: Bukan Palindrom")	
	}
}