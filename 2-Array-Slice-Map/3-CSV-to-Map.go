package main

import (
	"fmt"
	"strings"
)

func main() {

	var keys = "firstName,lastName,nationality"
	var values = "Sergei,Dragunov,Russia"

	splittedKeys := strings.Split(keys, ",")
	splittedValues := strings.Split(values, ",")
	temp := map[string]string{}

	for i := 0; i < len(splittedKeys); i++ {
		temp[splittedKeys[i]] = splittedValues[i]
	}
	fmt.Println(temp)
}
