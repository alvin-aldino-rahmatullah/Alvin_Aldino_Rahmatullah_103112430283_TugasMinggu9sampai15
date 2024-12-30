package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Input harus bilangan positif")
		return
	}

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {

			if i == j || i+j == n-1 {
				fmt.Print(j+1, " ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
