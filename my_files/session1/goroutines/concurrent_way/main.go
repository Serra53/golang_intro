package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go isUp(l, c)
	}

	for l := range c {
		//annonymous function
		go func(link string) {
			time.Sleep(10 * time.Second)
			isUp(link, c)
		}(l)

	}

}

func isUp(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error in Link %v: %v\n", link, err)
		c <- link
		return
	}
	fmt.Printf("%v is %v\n", link, resp.Status)
	c <- link
}
