package main

import "fmt"

func main() {
	var a, c int
	a, c = rectangle(20, 30)
	fmt.Println("area:", a)
	fmt.Println("parameter:", c)
}
func rectangle(l int, b int) (area int, parmeter int) {
	parmeter = 2 * (l + b)
	area = l * b
	return
}
