package main

import "fmt"

func main() {
    var huruf rune

    fmt.Print("Masukkan satu karakter: ")
    fmt.Scan(&huruf)

    if (huruf >= 'a' && huruf <= 'z') || (huruf >= 'A' && huruf <= 'Z') {
        if huruf != 'a' && huruf != 'e' && huruf != 'i' && huruf != 'o' && huruf != 'u' &&
            huruf != 'A' && huruf != 'E' && huruf != 'I' && huruf != 'O' && huruf != 'U' {
            fmt.Println("konsonan")
        } else {
            fmt.Println("bukan konsonan")
        }
    } else {
        fmt.Println("bukan konsonan")
    }
}
