package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//io.Copy Copies from a src to a destination
	//in this case resp.Body is the src Reader
	// and dst is os.Stdout (which is a simple print to the terminal)
	// the parameters are two types that implement their respective interfaces
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just worte this many bytes: ", len(bs))
	return len(bs), nil
}
