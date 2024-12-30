package main

import "fmt"

func main() {
    var n, m, x, y, cangkir int

    fmt.Print("Masukkan jumlah gula, kopi, kebutuhan gula per cangkir, kebutuhan kopi per cangkir (pisahkan dengan spasi): ")
    fmt.Scan(&n, &m, &x, &y)

    for n >= x && m >= y {
        n -= x  
        m -= y  
        cangkir++ 
    }

    fmt.Println("Jumlah cangkir kopi yang bisa dibuat:", cangkir)
}
