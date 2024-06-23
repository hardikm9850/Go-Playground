package main

// Deeper into Go
// https://www.udemy.com/course/go-the-complete-developers-guide/
func main() {
	cards := createDeck()
	cards.print()

	println("========Shuffling the deck ========\n\n")
	cards.shuffle()
	cards.print()

	println("========Dealing with the deck ========\n\n")
	dealWithCards(cards)

	println("========Saving the deck to storage ========\n\n")
	saveDeckToStorage(cards)

	println("========Creating the deck from the storage ========\n\n")
	createDeckFromStorage()
}

func createDeckFromStorage() {
	deck := createDeckFromFile("saved_deck")
	deck.print()
}

func saveDeckToStorage(cards deck) {
	fileSaveError := cards.saveDeckToFile("saved_deck")
	if fileSaveError != nil {
		panic(fileSaveError)
	}
}

func dealWithCards(cards deck) {
	hand, remainingCards := deal(cards, 3)
	println("* cards in hand*")
	hand.print()
	println("=====")
	println("* cards left *")
	remainingCards.print()
}

func createDeck() deck {
	cards := newDeck()
	return cards
}
