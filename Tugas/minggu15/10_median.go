package main

import "fmt"

func main() {
    var j, y, bilangan int
    total := 0

    fmt.Print("Masukan bilangan Y: ")
    fmt.Scan(&y)

    fmt.Println("Masukan 9 Bilangan: ")
    for j = 1; j <= 9; j++ {
        fmt.Scan(&bilangan)
        total += bilangan
    }


    if total >= y*5 {
        fmt.Println("Median bernilai", y)
    } else {
        fmt.Println("Median bernilai: 0")
    }
}

