/*Write a C++ program to swap first and last digits of any number. Go to the editor
Sample Output:
Input any number: 12345
The number after swapping the first and last digits are: 52341*/
package main

import "fmt"

func main() {
	var num, dig, result int

	fmt.Println("enter number")
	fmt.Scanf("%d", &num)

	for num != 0 {
		dig = num % 10
		num = num / 10
		result = dig
	}
	fmt.Print("Result = ", result)
}
