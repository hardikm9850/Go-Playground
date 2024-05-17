package main

// Deeper into Go
// https://www.udemy.com/course/go-the-complete-developers-guide/
func main() {
	//cards := deck{"Ace of A", "Ace of B", "Ace of C"}
	//cards = append(cards, "Ace of D")

	cards := newDeck()

	cards.print()
	// for i := 0; i < len(cards); i++ {
	//println(i, cards[i])
	// }
}
