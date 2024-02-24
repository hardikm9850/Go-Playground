package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"hardik", "payal"}
	message, error := greetings.Hellos(names)
	if error != nil {
		log.Fatal(error)
		return
	}

	for name, greeting := range message {
		fmt.Printf("Hey %s, You received %s\n", name, greeting)
	}
}
