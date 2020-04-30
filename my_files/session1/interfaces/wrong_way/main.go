package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreetingE(eb)
	printGreetingS(sb)

}

func printGreetingE(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingS(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
