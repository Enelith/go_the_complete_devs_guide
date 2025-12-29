package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades") // Add an element to the end of the slice

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
