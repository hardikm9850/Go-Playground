package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// handSize = 3
	// d[:3] will return cards from index 0 to 2, that is to say it's non-inclusive
	// d[3:] will return all cards from index 3
	return d[:handSize], d[handSize:]
}

func (d deck) stringFormat() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.stringFormat()), 0666)
}
