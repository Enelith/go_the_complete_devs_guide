package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://humblebundle.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// Creating a Channel of type string, and it's scoped inside the main function only
	c := make(chan string)

	for _, link := range links {
		// Passing the channel to the goroutine as an argument to the function checkLink
		go checkLink(link, c)
	}

	fmt.Println("Waiting for all goroutines to finish...")
	// This is a Blocking Call / Blocking Channel as it's waiting for the 1st msg to be retrieved through the Channel (and won't wait for another to finish)
	// In order to get all the msg, we would need as many fmt.Println as the links Slice length, so we opt for a loop.
	for _ = range links {
		fmt.Println(<-c)
	}
	/*
		for i := 0; i < len(links); i++ {
			fmt.Println(<-c)
		}
	*/
}

func checkLink(link string, c chan string) {
	// This line is called a "Blocking Call" as it will freeze the all programs until it gets an answer
	response, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, " might be down")
		//fmt.Println("Error: ", err)
		c <- fmt.Sprintf("[Goroutine & Channels] %v might be down!", link)
		return
	}

	//fmt.Println(link, " is up")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	c <- fmt.Sprintf("[Goroutine & Channels] %v is up!", link)
}
