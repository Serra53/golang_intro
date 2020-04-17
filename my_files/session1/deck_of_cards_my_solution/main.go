package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Shuffling new deck")
	deck := newDeck()
	fmt.Println(deck)
	shuffledDeck := shuffle(deck)
	fmt.Println("----------")
	print(shuffledDeck)
	deltCards := deal(shuffledDeck, 5)
	fmt.Println(deltCards)

}

func newDeck() []string {
	//this is a function that creates a new deck
	FullDeck := []string{}
	listOfNaipes := []string{"D", "C", "S", "H"}
	cards := []string{
		"A", "2", "3", "4", "5", "6", "7", "8", "9",
		"10", "J", "Q", "K", "A"}
	for i := 0; i < len(listOfNaipes); i++ {
		for j := 0; j < len(cards); j++ {
			card := listOfNaipes[i] + cards[j]
			FullDeck = append(FullDeck, card)
		}
	}
	return FullDeck
}

func shuffle(deck []string) []string {
	fmt.Printf("printing the deck of %v", len(deck))
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func print(deck []string) {
	for i := range deck {
		fmt.Println(deck[i])
	}
}

func deal(deck []string, numberPlayers int) [][]string {
	cardsDelt := [][]string{}
	if len(deck)%numberPlayers == 0 {
		numberCards := len(deck) / numberPlayers
		for i := 1; i <= numberPlayers; i++ {
			cards := deck[:numberCards]
			fmt.Printf("-------Cards Delt - Player %v ------------", i)
			fmt.Println(cards)
			cardsDelt = append(cardsDelt, cards)
			deck = deck[numberCards:]
		}
	} else {
		deck = deck[len(deck)%numberPlayers:]
		numberCards := len(deck) / numberPlayers
		for i := 1; i <= numberPlayers; i++ {
			cards := deck[:numberCards]
			fmt.Printf("-------Cards Delt - Player %v ------------", i)
			fmt.Println(cards)
			cardsDelt = append(cardsDelt, cards)
			deck = deck[numberCards:]
		}
	}
	return cardsDelt
}
