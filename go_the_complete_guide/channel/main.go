package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://msn.com",
		"http://github.com",
		"http://quora.com",
	}

	// create channel for inter-communication between goroutines
	channel := make(chan string)

	for _, link := range links {
		go checkLink(channel, link)
	}

	// Print channel data
	//for i := 0; i < len(links); i++ {
	// fmt.Println(<-channel)
	//}

	// Hit website checker forever
	for chanelRange := range channel {
		//go checkLink(channel, chanelRange)
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(channel, chanelRange)
		}(chanelRange)
	}
}

func checkLink(channel chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		channel <- link
		return
	}

	fmt.Println(link, "is up!")
	channel <- link
}
