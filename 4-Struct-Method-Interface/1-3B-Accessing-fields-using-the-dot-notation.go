package main

import (
	"fmt"
)

type Customer struct {
	Id      int
	Name    string
	addr    address
	married bool
}

type address struct {
	Number  int
	City    string
	Country string
}

func main() {
	var newCustomer = Customer{
		Id:   1,
		Name: "John",
		addr: address{
			Number:  20,
			City:    "Jakarta",
			Country: "Indonesia",
		},
		married: true,
	}

	var marr string
	if newCustomer.married == true {
		marr = "married"
	} else {
		marr = "not married"
	}

	fmt.Printf("Hello, My Name is %s\n", newCustomer.Name)
	fmt.Printf("I live in %s %s on Jl. Setia Budi No %d\n", newCustomer.addr.City, newCustomer.addr.Country, newCustomer.addr.Number)
	fmt.Printf("I am %s man\n", marr)
}
