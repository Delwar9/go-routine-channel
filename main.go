package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Create the channels
	evenCh, oddCh := make(chan bool, 1), make(chan bool, 1)
	// Defer the closing of the channels
	defer close(evenCh)
	defer close(oddCh)

	// Create the wait group
	wg = sync.WaitGroup{}
	// Add the number of goroutines to the wait group
	wg.Add(2)

	// Start the goroutines
	go even(evenCh, oddCh)
	go odd(oddCh, evenCh)

	// Start the sequence
	evenCh <- true

	// Wait for the goroutines to finish
	wg.Wait()
}

// even prints the even numbers

func even(evenCh, oddCh chan bool) {
	for i := 2; i <= 20; i += 2 {
		// Wait for the odd goroutine to signal
		<-oddCh
		fmt.Println(i)
		// Signal the odd goroutine
		evenCh <- true
	}

	// Signal the wait group that this goroutine is done
	wg.Done()
}

// odd prints the odd numbers

func odd(oddCh, evenCh chan bool) {
	for i := 1; i <= 20; i += 2 {
		// Wait for the even goroutine to signal
		<-evenCh
		fmt.Println(i)
		// Signal the even goroutine
		oddCh <- true
	}

	// Signal the wait group that this goroutine is done
	wg.Done()
}
