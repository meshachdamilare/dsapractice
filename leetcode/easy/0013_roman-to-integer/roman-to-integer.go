package main

import "fmt"

func main() {
	fmt.Println(RomanToInt("CX"))
}

func RomanToInt(s string) int {
	ans := 0
	num := 0
	var i int
	for i = len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case 'I':
			num = 1
		case 'V':
			num = 5
		case 'X':
			num = 10
		case 'L':
			num = 50
		case 'C':
			num = 100
		case 'D':
			num = 500
		case 'M':
			num = 1000
		}
		if 4*num < ans {
			ans -= num
		} else {
			ans += num
		}
	}
	return ans
}
