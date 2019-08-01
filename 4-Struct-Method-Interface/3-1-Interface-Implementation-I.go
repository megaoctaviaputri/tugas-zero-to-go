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

func main() {
	var NewUser Namer
	NewUser = &User{"Matt", "Aimonetti"}
	fmt.Println(NewUser.Name())
}
