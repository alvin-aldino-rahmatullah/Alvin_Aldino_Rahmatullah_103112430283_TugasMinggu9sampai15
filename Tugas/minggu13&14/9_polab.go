package main 
import "fmt"

func main() {
	var k, j, n int

	fmt.Print("Masukkan bilangan positif : ")
	fmt.Scan(&n)

	for j = 1; j <= n; j++ {
		for  k = 1; k <= n; k++ {
			if j == 1 || j == n || k == 1 || k == n {
				fmt.Print(j," ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}