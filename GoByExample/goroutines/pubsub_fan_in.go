package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")

	// Create a buffered channel with a buffer size of 10
	channel := make(chan int, 10)
	errorChan := make(chan error)

	// Create multiple publishers
	for i := 0; i < 5; i++ {
		go func(publisherID int) {
			for j := 1; j <= 2; j++ {
				fmt.Printf("Publisher %d Publishing: %d\n", publisherID, j)
				channel <- j
				time.Sleep(time.Millisecond * 500) // adding a delay
			}
		}(i)
	}

	// Create a single subscriber
	subscribeFanIn(channel, errorChan)

	fmt.Println("The End")
}

func subscribeFanIn(channel chan int, errorChan chan error) {
	timeout := time.After(5 * time.Second)
	rateLimiter := time.Tick(time.Second * 2) 

	for {
		select {
		case value, ok := <-channel:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			<-rateLimiter // Preventing rate limiting 
			// it ensures that we wait for the next value from the rateLimiter channel before processing each message
			fmt.Printf("Received: %d\n", value)

		case err := <-errorChan: //Error Handling:
			fmt.Printf("Error received: %v\n", err)
			return
		case <-timeout:
			fmt.Println("Timeout received")
		default:
		}
	}
}

// Rate Limiting
