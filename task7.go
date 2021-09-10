package main

import "fmt"

func main() {
	fmt.Println("enter first number")
	var p1 int
	fmt.Scanln(&p1)
	fmt.Println("enter second number")
	var p2 int
	fmt.Scanln(&p2)
	fmt.Println("enter third number")
	var p3 int
	fmt.Scanln(&p3)
	fmt.Println("enter forth number")
	var p4 int
	fmt.Scanln(&p4)
	fmt.Println("enter fifth number")
	var p5 int
	fmt.Scanln(&p5)
	sum := p1 + p2 + p3 + p4 + p5
	avg := sum / 5
	fmt.Println("averag:", avg)
}
