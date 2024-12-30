package main 
import "fmt"

func main() {
	var k, j, n int

	fmt.Print("Masukkan bilangan positif : ")
	fmt.Scan(&n)

	for k = 1; k <= n; k++ {
		for  j = 1; j <= n; j++ {
			fmt.Print(j," ")
		}
		fmt.Println()
	}

}