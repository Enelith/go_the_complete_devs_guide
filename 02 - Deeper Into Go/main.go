package main

func main() {
	cards := newDeck()
	_ = cards.saveToFile("deck.txt")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
