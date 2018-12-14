package main

import (
	"log"
	"sync"
	"time"
)

func inputLoop() {
	for i := 0; i < 10; i++ {
		log.Printf("Retrieving %v", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func displayLoop() {
	for i := 0; i < 10; i++ {
		log.Printf("Displaying %v", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// Main wait group
	var wg sync.WaitGroup

	// Create our input go routine and add to our wait group
	wg.Add(1)
	go func() {
		inputLoop()
		wg.Done()
	}()

	// Create our display go routine and add to wait group
	wg.Add(1)
	go func() {
		displayLoop()
		wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
