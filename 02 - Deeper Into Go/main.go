package main

import "fmt"

const filename = "deck.txt"

func main() {
	cards := newDeck()
	fmt.Println("Create new deck: ", cards)
	fmt.Println("Saving to file: ", filename)
	_ = cards.saveToFile(filename)

	cards = nil
	fmt.Println("Reset deck: ", cards)

	fmt.Println("Loading from file: ", filename)
	cards = newDeckFromFile(filename)
	fmt.Println("Loaded deck: ", cards)
	//	cards.print()

	fmt.Println("Deal 5 cards: ")
	hand, remainingCards := deal(cards, 5)

	fmt.Println("Hand: ", hand)
	//	hand.print()

	fmt.Println("Remaining cards: ", remainingCards)
	//	remainingCards.print()
}
