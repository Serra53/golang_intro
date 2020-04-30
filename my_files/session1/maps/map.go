package main

import "fmt"

func main() {
	//var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "fffffff",
	}

	fmt.Println(colors)
	fmt.Printf("%+v\n", colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex Code for %v is %v.\n", color, hex)
	}
}
