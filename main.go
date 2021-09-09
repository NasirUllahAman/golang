package main

import (
	"fmt"
)

func main() {

	var slice1 []int
	var slice2 []int

	i := 0
	for i <= 1000 {
		if i%2 == 0 {
			slice1 = append(slice1, i)
		} else {
			slice2 = append(slice2, i)
		}
		i++
	}
	counter := 0
	c := 0
	for j := 2; j <= 1000; j++ {
		for k := 0; k < i; k++ {
			counter++

		}
		if counter == j-2 {

			fmt.Println("prime number: %d\n", j)
			c++
		}
		counter = 0

	}

	fmt.Println("Even number:", slice1)
	fmt.Println("Odd number", slice2)

}
