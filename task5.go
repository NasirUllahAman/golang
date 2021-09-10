package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}
	m := make(map[string]int)
	for i := 0; i < len(a); i++ {
		m[b[i]] = a[i]
	}
	fmt.Println(m)

}
