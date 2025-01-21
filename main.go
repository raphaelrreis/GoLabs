package main

import (
	"fmt"
	"math/rand"
	"time"
)

// hello simulates a task by printing a message and pausing for a random duration.
func hello(msg string) {
	fmt.Println(msg + " - goroutine")                        // Print the message with an identifier for the goroutine
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Second) // Simulate work with a random delay
}

func main() {
	// Start two goroutines that execute the hello function concurrently
	go hello("Hello 1")
	go hello("Hello 2")

	// Ensure the main function waits before exiting to allow goroutines to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Common Callback") // Final message printed after goroutines start
}
