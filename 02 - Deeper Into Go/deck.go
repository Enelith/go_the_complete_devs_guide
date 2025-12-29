package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	// Are slices of String
	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}
	return cards
}

func (d deck) print() {
	for index, card := range d {
		println(index, card)
	}
}

// Return the "hand" (from 0 to handSize of the original deck), and the rest of the deck (from handSize to the rest of the dec)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
