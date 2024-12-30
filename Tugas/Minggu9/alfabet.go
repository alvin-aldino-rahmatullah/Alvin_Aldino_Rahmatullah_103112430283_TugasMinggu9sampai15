package main 
import "fmt"

func main (){
	var char rune
	fmt.Scan(&char)
    
	if char <= 'a' &&  char >= 'z' || char <= 'A' && char >= 'Z' {
		fmt.Println("Alfabet")
	} else {
		fmt.Println("Bukan Alfabet")
	}
}