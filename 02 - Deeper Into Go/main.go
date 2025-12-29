package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades") // Add an element to the end of the slice

	for key, value := range cards {
		fmt.Println(key, value)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
