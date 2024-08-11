package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")

	channel := make(chan int)
	done := make(chan struct{})

	go publish(channel)
	go subscribe(channel, done)
	<-done // A hack to keep the main goroutine running

	fmt.Println("The End")
}

func subscribe(channel chan int, done chan struct{}) {
	for value := range channel {
		fmt.Printf("Received: %d\n", value)
	}
	close(done)
}

func publish(channel chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Publishing: %d\n", i)
		channel <- i
		time.Sleep(time.Millisecond * 500) // adding a delay
	}
	// This built-in function sets a flag indicating that no additional data will be sent to this channel
	// Leaving a channel open can cause memory leaks https://stackoverflow.com/a/74771266/2070800
	close(channel)
}
