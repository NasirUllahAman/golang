package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}
	var list3 []string
	j := 0
	k := 0
	for i := 0; i < len(a)+len(b); i++ {
		if i%2 == 0 {
			list3 = append(list3, b[j])
			j++
		} else {
			list3 = append(list3, string(a[k]))
			k++
		}
	}
	fmt.Println(list3)
}
