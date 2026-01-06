package main

import (
	"fmt"
	"io"
	"net/http"
)

// Result display as follow (same order as defined) with a delay between each check:
// http://google.com  is up
// http://facebook.com  is up
// http://humblebundle.com  is up
// http://stackoverflow.com  is up
// http://golang.org  is up
// http://amazon.com  is up
func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://humblebundle.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	// This line is called a "Blocking Call" as it will freeze the all programs until it gets an answer
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(link, " is up")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
}
