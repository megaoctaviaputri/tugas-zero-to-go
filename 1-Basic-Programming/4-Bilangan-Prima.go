package main

import "fmt"

func main() {
        var bilangan int
        var count int

        count = 0

        fmt.Scan(&bilangan)

        for i := 1; i <= bilangan; i++ {
                if bilangan%i == 0 {
                        count += 1
                }
        }
        if count == 2 {
                fmt.Println("Bilangan Prima")
        } else {
                fmt.Println("Bukan Bilangan Prima")
        }

}