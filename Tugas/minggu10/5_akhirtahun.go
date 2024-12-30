package main

import "fmt"

func main() {
	var totalBelanja, diskon, cashback int
	var buatKartu bool

	fmt.Print("Masukkan total belanja : ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Apakah mau membuat kartu? (true/false): ")
	fmt.Scan(&buatKartu)

	if buatKartu {
		fmt.Println("Kartu? true")
	} else {
		fmt.Println("Kartu? false")
	}

	if totalBelanja >= 100000 {
		fmt.Println("Diskon? true")
		diskon = totalBelanja * 10 / 100
		totalBelanja -= diskon
	} else {
		fmt.Println("Diskon? false")
	}

	if totalBelanja >= 200000 {
		fmt.Println("Cashback? true")
		cashback = 75000
		totalBelanja -= cashback
	} else {
		fmt.Println("Cashback? false")
	}

	fmt.Println("Rp.", totalBelanja)
}
