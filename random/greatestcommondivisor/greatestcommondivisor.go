package main

import "fmt"

func main() {
	fmt.Println(GCD(8, 64))
	fmt.Println(3 % 4)
}

func GCD(a, b int) int {
	for b != 0 {
		rem := a % b
		a = b
		b = rem

	}
	return a
}
