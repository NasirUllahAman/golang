package main

import "fmt"

func main() {
	//i := 0
	//num := 0
	primeNum := 0

	for i := 2; i <= 1000; i++ {
		counter := 0
		for num := 2; num <= i/2; num++ {
			if i%num == 0 {
				counter = counter + 1
				break

			}
		}
		if counter == 0 {
			fmt.Print(" ", i)
			primeNum = primeNum + 1

		}
	}
	fmt.Println("\nTotal Prime Numbers:", primeNum)

}
