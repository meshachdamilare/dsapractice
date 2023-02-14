package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	court := make(chan int)

	// To manage our goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		player("PlayerA", court)
		wg.Done()
	}()

	go func() {
		player("PlayerB", court)
		wg.Done()
	}()

	// After guarantee of a receiver, It sends signal to start the game
	court <- 1

	// Wait for the game to finish
	wg.Wait()
}

func player(name string, court chan int) {
	for {

		// Wait for the ball to be hit back to us.
		ball, wd := <-court
		if !wd {

			// If channel is closed, nothing was sent back; WE WON
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Trick: Pick a random number and see if we miss the ball.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hits %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
