// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffer channel so no send every block. Don't allocate more buffers than
// you need. Have the main goroutine display each random number is receives and
// then terminate the program.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ch := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			ch <- rand.Intn(100)

		}()
	}

	waitRes := goroutines
	var nums []int
	for waitRes > 0 {
		nums = append(nums, <-ch)
		waitRes--
	}
	fmt.Printf("Values: %v\n", nums)
}
