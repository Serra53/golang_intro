package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	valueExpected := 52
	if len(d) != valueExpected {
		t.Errorf("Expected deck length of %v but got %v",
			valueExpected, len(d))
	}
}

func TestCardInDeck(t *testing.T) {
	card := "Four of Diamonds"
	d := newDeck()
	cardInDeck := false
	for _, item := range d {
		if item == card {
			cardInDeck = true
		}
	}
	if cardInDeck != true {
		t.Errorf("Card %v not found in new deck", card)
	}
}

func TestSaveDeckAndLoad(t *testing.T) {
	fp := "_decktesting"
	os.Remove(fp)
	deck := newDeck()
	deck.saveToFile(fp)
	fmt.Println(len(deck))
	loadedDeck := newDeckFromFile(fp)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards.")
	}
	os.Remove(fp)
}
