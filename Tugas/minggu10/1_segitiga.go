package main

import "fmt"

func main() {
	var a, b, c int
	
	fmt.Print("Masukkan Jumlah sisi segitiga : ")
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Print("Segitiga sama sisi")
	} else if a != b && b != c && a != c {
		fmt.Print("Segitiga sembarang")
	} else if a != b && b != c && a == c {
		fmt.Print("Segitiga sama kaki")
	}
}
