package main

import "fmt"

func main() {
	var input, input2 string         
	const username = "admin" 
	const password = "admin" 
	var attempt int = 0

	for {
		fmt.Print("Masukkan username dan password : ")
		fmt.Scan(&input, &input2) 

		attempt++

		if input == username && input2 == password {
			fmt.Println("Jumlah percobaan: ", attempt)
			fmt.Println("Login berhasil!")
			return
		} else {
			fmt.Println("Password salah.")
		}
	}
}