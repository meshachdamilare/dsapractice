package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindPrimeFactors(90))
	fmt.Println(FindPrimeFactors2(90))
}

/*
	Solution 1
*/

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

/*
	Solution 2
	This solution only gives the likely prime factors without repetition
*/

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrimeFactors2(n int) []int {
	var factor []int
	for i := 2; i <= n; i++ {
		if n%i == 0 && isPrime(i) {
			factor = append(factor, i)
		}
	}
	return factor
}
