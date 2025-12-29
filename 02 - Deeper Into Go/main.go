package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" // New Variable
	//card = "Five of Diamonds" // Variable assignment
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
