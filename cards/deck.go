package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	time2 "time"
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

func (d deck) saveDeckToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.stringFormat()), 0666)
}

func createDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	deckInString := string(data)
	return strings.Split(deckInString, ",")
}

func (d deck) shuffle() {
	time := time2.Now()
	source := rand.NewSource(time.UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
