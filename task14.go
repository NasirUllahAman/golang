package main

import "fmt"

func pnum(minNum int, c chan int) {
	primeNum := 0
	maxNum := minNum + 1999
	for i := minNum; i <= maxNum; i++ {
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
	c <- primeNum
}

func main() {

	//primeNum := 0
	var count, sum int
	c := make(chan int)
	go pnum(2, c)
	count = <-c
	sum += count
	fmt.Println("  ", count)
	go pnum(2000, c)
	count = <-c
	sum += count
	fmt.Println("  ", count)
	go pnum(4000, c)
	count = <-c
	sum += count
	fmt.Println("  ", count)
	go pnum(6000, c)
	count = <-c
	sum += count
	fmt.Println("  ", count)
	go pnum(8000, c)
	count = <-c
	fmt.Println("  ", count)
	sum += count

	fmt.Println("\nTotal Prime Numbers:", sum)

}
