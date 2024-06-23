package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card od Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveDeckToFile("_decktesting")
	loadedDeck := createDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
