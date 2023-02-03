package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindFactors(6))
}

func FindFactors(number int) []int {
	var factors []int
	for number%2 == 0 {
		factors = append(factors, 2)
		number = number / 2
	}
	i := 3

	max_factor := int(math.Sqrt(float64(number)))
	for i <= max_factor {
		for number%i == 0 {
			factors = append(factors, i)
			number = number / i
			max_factor = int(math.Sqrt(float64(number)))
		}
		i = i + 2
	}
	if number > 1 {
		factors = append(factors, number)
	}
	return factors
}
