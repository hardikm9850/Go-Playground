package main

import "fmt"

type englishBot struct {
}
type spanishBot struct {
}
type bot interface {
	getGreeting() string
}

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/**
 * Explanation : Since getGreeting method is created on both englishBot and spanishBot, they are now member of type 'bot'
 * Interfaces are 'implicit' in Go. This is very different from what I've seen in java/kotlin.
 * With this, we can call printGreeting method with above-given 2 instances
 */
func (e englishBot) getGreeting() string {
	return "Hello"
}

func (e spanishBot) getGreeting() string {
	return "Hola"
}
