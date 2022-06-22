package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Two of Hearts" {
		t.Errorf("Expected the first String to be Two of Hearts %v", d[0])
	}

	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected the first String to be Ace of Spades %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	err := os.Remove("_decktesting")
	if err != nil {
		return
	}

	deck := newDeck()
	err = deck.saveToFile("_decktesting")
	if err != nil {
		return
	}

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 200 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	err = os.Remove("_decktesting")
	if err != nil {
		return
	}
}
