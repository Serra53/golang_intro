package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, l := range links {
		isUp(l)
	}
}

func isUp(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error in Link %v: %v\n", link, err)
	}
	fmt.Printf("%v is %v\n", link, resp.Status)
}
