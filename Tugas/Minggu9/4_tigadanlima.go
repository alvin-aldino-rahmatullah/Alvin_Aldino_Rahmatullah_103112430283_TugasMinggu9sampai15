package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat : ")
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("kelipatan 3")
		fmt.Println("kelipatan 5")
	} else if n%5 == 0 {
		fmt.Print("kelipatan 5")
	} else if n%3 == 0 {
		fmt.Print("kelipatan 3")
	} else {
		fmt.Print("")
	}
}
