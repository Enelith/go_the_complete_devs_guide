package main

import (
	"fmt"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	//return strings.Join([]string(d), ",")
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return an empty deck (newDeck())
		// Option #2 - log the error and entirely quit the program (os.Exit(x > 0), panic)
		fmt.Println("Error reading file: ", err)
		//panic(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}
