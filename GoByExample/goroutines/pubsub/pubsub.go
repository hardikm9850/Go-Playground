package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")

	channel := make(chan int)
	errorChan := make(chan error)

	// goroutines are created using the `go`` keyword followed by a function call.
	go publish(channel, errorChan)
	subscribe(channel, errorChan)

	fmt.Println("The End")
}

func subscribe(channel chan int, errorChan chan error) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case value, ok := <-channel:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Received: %d\n", value)

		case err := <-errorChan: //Error Handling:
			fmt.Printf("Error received: %v\n", err)
			return
		case <-timeout:
			fmt.Println("Timeout received")
		}
	}
}

func publish(channel chan int, errorChan chan error) {
	for i := 1; i <= 14; i++ {
		fmt.Printf("Publishing: %d\n", i)
		channel <- i
		time.Sleep(time.Millisecond * 500) // adding a delay
		
		// Cause an error
		/* if i == 5 {
			errorChan <- fmt.Errorf("Error occurred during publishing")
			return
		} */

	}
	// This built-in function sets a flag indicating that no additional data will be sent to this channel
	// Leaving a channel open can cause memory leaks https://stackoverflow.com/a/74771266/2070800
	close(channel)
	close(errorChan)
}

// Timeouts
// Buffered Channels
// Fan-In/Fan-Out Pattern
// Graceful Shutdown:
// Rate Limiting
