package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
		//panic(err)
	}

	// 1) Using the READ method
	// Creating an empty byte slice large enough to hold the response body (if too small, the Reader will stop early)
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)

	// Will therefore display the response body html
	//fmt.Println(string(bs))

	// 2) Using the COPY method
	//io.Copy(os.Stdout, resp.Body)

	// 3) Using a Custom Writer
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))

	// Returning the length of the byte slice we consumed through the fmt.Println
	return len(bs), nil
}
