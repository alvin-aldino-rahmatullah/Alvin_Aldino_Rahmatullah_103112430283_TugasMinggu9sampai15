package main 
import "fmt"

func main() {
	var num, counter int
	fmt.Scan(&num)

    counter = 0

	for num > 0 {
       if num % 10 % 2 == 0 {
		counter = counter + 1
	   }
	   num = num / 10
	}

	fmt.Print(counter)
}