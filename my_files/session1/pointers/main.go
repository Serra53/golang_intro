package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
