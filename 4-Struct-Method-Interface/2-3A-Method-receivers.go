package main

import "fmt"

type User struct {
	FirstName, LastName string
}

// merupakan method pointer
// ketika dipanggil maka akan merubah isi variabel dari struct User sesuai dengan apa yang diperintahkan oleh method ini
// karena yang di-copy oleh method ini adalah alamat dari struct user
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"Matt", "Aimonetti"} // variabel u menampung alamat dari struct User
	fmt.Println(u.Greeting())
}
