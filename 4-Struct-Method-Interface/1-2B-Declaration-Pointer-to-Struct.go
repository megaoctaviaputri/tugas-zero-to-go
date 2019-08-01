package main

import (
	"fmt"
)

type Student struct {
	Rollumber int
	Name      string
}

func main() {
	// do something here
	s := Student{11, "Jack"}
	ps := &s
	fmt.Println(ps)
	fmt.Println(ps.Name)
	fmt.Println(s.Name)
	ps.Rollumber = 31
	fmt.Println(ps)
}
