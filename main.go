package main

func main() {
	cards := newDeck()
	cards.shuffleDeck()
	cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards.saveToFile("my_cards")

}
