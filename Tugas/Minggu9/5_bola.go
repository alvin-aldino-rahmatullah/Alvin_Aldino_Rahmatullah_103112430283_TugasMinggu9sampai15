package main

import "fmt"

func main() {
    var a, b, c, d, max, min int

    fmt.Print("Masukkan jumlah gol empat tim : ")
    fmt.Scan(&a, &b, &c, &d)

    if a > b {
        max = a
    } else {
        max = b
    }
    if c > max {
        max = c
    }
    if d > max {
        max = d
    }

    if a < b {
        min = a
    } else {
        min = b
    }
    if c < min {
        min = c
    }
    if d < min {
        min = d
    }

    fmt.Println("Gol terbanyak:", max)
    fmt.Println("Gol tersedikit:", min)
}
