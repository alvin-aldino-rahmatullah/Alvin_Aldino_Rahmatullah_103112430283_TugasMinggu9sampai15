package main

import "fmt"

func main() {
    var uang, total float64
	var assemen bool

	fmt.Print("Masukkan uang kamu : ")
	fmt.Scan(&uang)
	fmt.Print("Apakah kamu sedang assasmen (true/false) ? : ")
	fmt.Scan(&assemen)

	if assemen == true {
		total = uang * 0.65 
		fmt.Print("Kamu memperoleh potongan 35% dari ", uang, " menjadi ", total)
	} else {
       fmt.Println("Kamu tidak sedang assesmen jadi tidak ada diskon")
	}
	
}
