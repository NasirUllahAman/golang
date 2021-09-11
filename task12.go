package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("em4:=", a)
	fmt.Println("digit=", a[4])
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("lenght=", b)
	var twod [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println("twod=", twod)

}
