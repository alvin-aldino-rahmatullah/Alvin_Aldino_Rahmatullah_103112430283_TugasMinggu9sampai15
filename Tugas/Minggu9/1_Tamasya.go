package main

import "fmt"

func main() {
	var n, mobil int
	fmt.Print("berapa orang yang akan naik mobil ? : ")
	fmt.Scan(&n)

	if n % 7 == 0 {
		mobil = n / 7
	} else {
		mobil = n / 7 + 1
	}
	fmt.Print("jumlah mobil anda adalah ", mobil)

}
