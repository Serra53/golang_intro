package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	hand := cards.deal(5)
	cards.print()
	fmt.Println("Hand delt")
	hand.print()

}
