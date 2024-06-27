/**
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	mapOfWords := make(map[string]int)
	for _,v := range strings.Fields(s) {
		mapOfWords[v] = 1 + mapOfWords[v]
	}
	return mapOfWords
}

func main() {
	wc.Test(WordCount)
}
