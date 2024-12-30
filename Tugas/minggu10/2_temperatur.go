package main

import "fmt"

func main() {
	var a, b, c, d, e float64

	fmt.Print("Masukkan lima nilai temperatur: ")
	fmt.Scan(&a, &b, &c, &d, &e)

	if a > b && b > c && c > d && d > e {
		fmt.Print("stabil naik")
	} else if a < b && b < c && c < d && d < e {
		fmt.Print("stabil turun")
	} else if a != b && b != c && c != d && d != e {
		fmt.Print("tidak stabil")
	}
}

