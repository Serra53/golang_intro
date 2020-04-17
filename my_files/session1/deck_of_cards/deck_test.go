package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	valueExpected := 52
	if len(d) != valueExpected {
		t.Errorf("Expected deck length of %v but got %v", valueExpected, len(d))
	}
}

func TestCardInDeck(t *testing.T) {
	card := "Four of Diamonds"
	d := newDeck()
	cardInDeck := false
	for _, item := range d {
		if item == card {
			cardInDeck = true
			fmt.Println(item)
		}
	}
	if cardInDeck != true {
		t.Errorf("Card %v not found in new deck", card)
	}
}

// func TestSaveAndLoad() {

// }
