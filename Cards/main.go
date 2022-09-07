package main

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("my-cards")
	cards.shuffleDeck()
	cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("my-cards")
}

func newCard() string {
	return "King of Diamonds"
}