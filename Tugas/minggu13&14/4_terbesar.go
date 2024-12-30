package main

import "fmt"

func main() {
	var digit, maxDigit int

	fmt.Print("Masukkan bilangan positif : ")
	fmt.Scan(&digit)

	maxDigit = -1

	for digit > 0 {
		if digit%10 > maxDigit {
			maxDigit = digit % 10
		}
		digit = digit / 10
	}

	fmt.Print(maxDigit)
}
