package main

import (
	"fmt"
	"strings"
)

func main() {

	var inputString = "hello world"
	var arrayOfChar = []byte{'h', 'e', 'w', 'o'}

	for i := 0; i < len(arrayOfChar); i++ {
		inputString = strings.ReplaceAll(inputString, string(arrayOfChar[i]), "")
	}

	fmt.Println(inputString)

}
