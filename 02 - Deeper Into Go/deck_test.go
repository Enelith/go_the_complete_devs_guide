package main

import (
	"os"
	"testing"
)

const testFilename = "_decktesting"

// 2.36 Writing Useful Tests
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(d))
	}
}

// 2.37 Asserting Elements in a Slice
func TestAssertingElementInSlice(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card to be King of Spades, but got %v", d[len(d)-1])
	}
}

// 2.38 Testing File IO
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deleteTestFile()

	deck := newDeck()
	_ = deck.saveToFile(testFilename)

	loadedDeck := newDeckFromFile(testFilename)

	if len(loadedDeck) != len(deck) {
		t.Errorf("Loaded deck should have same length as original deck (%v != %v)", len(loadedDeck), len(deck))
	}

	deleteTestFile()
}

func deleteTestFile() {
	_ = os.Remove(testFilename)
}
