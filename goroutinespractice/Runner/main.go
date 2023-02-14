// Sample program to show how to use an unbuffered channel to
// simulate a relay race between four goroutines.

package main

import (
	"fmt"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {

	track := make(chan int)
	wg.Add(1)
	go Runner(track) // Gives guarantee of the receiver before sending
	track <- 1
	wg.Wait()
}

func Runner(track chan int) {
	const maxExchange = 4
	var exchange int

	baton := <-track

	fmt.Printf("Runner %d with Balton\n", baton)
	if baton < maxExchange {
		exchange = baton + 1
		fmt.Printf("Runner %d To the Line\n", exchange)

		// Here, track is empty and there's guarantee of a receiver but won't run until data/signal is sent in line:52
		// This can be place anywhere above line: 53
		go Runner(track)
	}

	time.Sleep(100 * time.Millisecond)

	if baton == maxExchange {
		fmt.Printf("Runner %d Finished, Race over", baton)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n",
		baton,
		exchange)

	// Data/Signal is sent to ensure thw worker on line:37 can start
	track <- exchange
}
