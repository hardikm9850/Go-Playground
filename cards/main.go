package main

// Deeper into Go
// https://www.udemy.com/course/go-the-complete-developers-guide/
func main() {
	cards := createDeck()

	dealWithCards(cards)

	fileSaveError := cards.saveToFile("saved_deck")
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
