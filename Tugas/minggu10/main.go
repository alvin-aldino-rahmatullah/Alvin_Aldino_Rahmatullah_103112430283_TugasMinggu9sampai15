package main

import "fmt"

func main() {
	var val int
	fmt.Print("masukkan nilai:")
	fmt.Scan(&val)

	if val >= 0 {
		fmt.Print("Nilai yang anda masukkan adalah positif")
		if val % 2 == 0  {
			
		}
	} else {
		fmt.Print("Nilai yang anda masukkan adalah negatif")
	}
}
