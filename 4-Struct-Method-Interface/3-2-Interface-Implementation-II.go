package main

import "fmt"

type Namer interface {
	Name() string
}

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

func main() {
	var NewUser Namer
	NewUser = &User{"Matt", "Aimonetti"}
	fmt.Println(NewUser.Name())

	var NewUser2 Namer
	NewUser2 = &Customer{42, "Francesc"}
	fmt.Println(NewUser2.Name())
}
