// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		valuePasser("A", ch)
		wg.Done()
	}()

	go func() {
		valuePasser("B", ch)
		wg.Done()
	}()

	ch <- 1

	wg.Wait()
}

func valuePasser(name string, share chan int) {
	for {
		value, ok := <-share
		if !ok {
			fmt.Printf("Goroutine %s is Down\n", name)
			return
		}

		fmt.Printf("Goroutine %s value %d\n", name, value)
		if value == 10 {
			close(share)
			fmt.Printf("Goroutine %s is Down\n", name)
			return
		}
		share <- (value + 1)
	}

}
