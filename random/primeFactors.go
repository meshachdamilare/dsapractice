package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindPrimeFactors(81))
}

func FindPrimeFactors(number int) []int {
	var factors []int
	for number%2 == 0 {
		factors = append(factors, 2)
		number = number / 2
	}
	i := 3

	maxFactor := int(math.Sqrt(float64(number)))
	for i <= maxFactor {
		for number%i == 0 {
			factors = append(factors, i)
			number = number / i
			maxFactor = int(math.Sqrt(float64(number)))
		}
		i = i + 2
	}
	if number > 1 {
		factors = append(factors, number)
	}
	return factors
}
