package main

import "fmt"

func main() {
	var x, n int

	fmt.Print("Masukkan bilangan x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan n : ")
	fmt.Scan(&n)

	found := false
	for n > 0 {
		digit := n % 10
		if digit == x {
			found = true
			break
		}
		n /= 10
	}

	if found {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
