package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//Create a new type of deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	newDeck := deck{}

	cardSuits := []string{"Spades", "Hearts",
		"Diamonds", "Clubs"}
	cardNumbers := []string{
		"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}

	for _, s := range cardSuits {
		for _, n := range cardNumbers {
			newCard := n + " of " + s
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	deckString := strings.Split(string(bs), "s")
	return deck(deckString)
}
