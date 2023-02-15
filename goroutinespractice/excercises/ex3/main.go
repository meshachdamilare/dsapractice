// Write a program that uses goroutines to generate up to 100 random numbers.
// Do not send values that are divisible by 2. Have the main goroutine receive
// values and add them to a slice.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	ch := make(chan int)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			n := rand.Intn(100)
			if n%2 == 0 {
				return
			}
			ch <- n
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var nums []int
	for val := range ch {
		nums = append(nums, val)
	}
	fmt.Printf("Value of slice: %v\n", nums)
}
