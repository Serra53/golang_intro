package main

import (
	"fmt"
	"math/rand"
)

//Create a new type of deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	newDeck := deck{}

	cardSuits := []string{"Spades", "Hearts",
		"Diamonds", "Clubs"}
	cardNumbers := []string{
		"Aces", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}

	for _, s := range cardSuits {
		for _, n := range cardNumbers {
			newCard := s + " of " + n
			newDeck = append(newDeck, newCard)
		}
	}
	return newDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(sizeOfHand int) deck {
	d = d[sizeOfHand:]
	return d[:sizeOfHand]
}

func (d deck) shuffle() deck {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}
