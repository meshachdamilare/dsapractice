// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	values := make(chan int)
	shutdown := make(chan struct{})

	poolSize := runtime.GOMAXPROCS(0)

	for i := 0; i < poolSize; i++ {
		go func(d int) {
			for {
				n := rand.Intn(100)
				select {
				case values <- n:
					fmt.Printf("Worker %d sent %d\n", d, n)
				case <-shutdown:
					fmt.Printf("Worker %d shutdown\n")
					wg.Done()
					return
				}
			}
		}(i)
	}

	var nums []int

	for val := range values {
		if val%2 == 0 {
			fmt.Printf("Discarding %d \n", val)
		}
		fmt.Printf("Keeping %d\n", val)
		nums = append(nums, val)
		if len(nums) == 100 {
			break
		}
	}
	fmt.Println("Receiver sending shutdown signal")
	close(shutdown)

	wg.Wait()
	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}
