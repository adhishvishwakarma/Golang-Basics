package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards := newDeckFromFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
