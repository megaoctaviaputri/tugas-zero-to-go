package main

import "fmt"

func palindrom(s string) bool {
	
	mapOfString := map[int]string{}
	var strMundur string

	for i := 0; i < len(s); i++ {
		mapOfString[i] = string(s[i])
	}

	for i := len(s) - 1; i >= 0; i-- {
		strMundur += mapOfString[i]
	}

	if s == strMundur {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(palindrom("katak"))
	fmt.Println(palindrom("blanket"))
	fmt.Println(palindrom("civic"))
	fmt.Println(palindrom("kasur rusak"))
	fmt.Println(palindrom("mister"))
}
