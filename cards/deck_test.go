package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	length := len(d)

	if length != 16 {
		t.Errorf("Expected deck length of 16, but got %d", length)
	}

	firstElem := d[0]
	if firstElem != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %s", firstElem)
	}

	lastEelem := d[length-1]
	if lastEelem != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs but got %s", lastEelem)
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	const path = "_decktesting"
	os.Remove(path)

	deck := NewDeck()
	deck.SaveToFile(path)

	loadedDeck := NewDeckFromFile(path)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedDeck))
	}

	os.Remove(path)
}
