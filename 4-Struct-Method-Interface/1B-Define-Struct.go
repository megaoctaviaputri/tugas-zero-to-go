package main

import (
	"fmt"
)

type Car struct {
	Name       string
	Model      string
	Color      string
	WeightInKg int
}

func main() {
	c := Car{
		Name:       "Ferrari",
		Model:      "GTC4",
		Color:      "Red",
		WeightInKg: 1920,
	}
	fmt.Println("Car: ", c)
}
