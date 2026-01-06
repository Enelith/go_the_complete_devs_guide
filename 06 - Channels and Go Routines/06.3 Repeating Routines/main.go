package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
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

	fmt.Println("Looping forever (and repeating each websites check after they've send their answer through the channel...)")
	/*
		for {
			// Using a Function Literals to add a delay between each check of Link.
			go func() {
				time.Sleep(time.Second * 2) // Sleep the process for 2 seconds
				// Go knows '<-c' will produce a string
				checkLink(<-c, c)
			}() // <- the '()' is added as it is a FUNCTION CALL
		}
	*/
	// Another way (easier to read) to write it :
	// "Watch the Channel 'c', whenever a value comes out of it, assign that value to 'passedLink'.
	// Once assigned, the body of the for loop is immediately executed 'go checkLink(passedLink, c)'"
	for passedLink := range c {
		// Using a Function Literals to add a delay between each check of Link.
		go func(argLink string) {
			time.Sleep(time.Second * 2) // Sleep the process for 2 seconds
			checkLink(argLink, c)
		}(passedLink) // <- the '()' is added as it is a FUNCTION CALL
	}
}

// We're going to add a delay to avoid spamming the check
func checkLink(link string, c chan string) {
	// This line is called a "Blocking Call" as it will freeze the all programs until it gets an answer
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		fmt.Println("Error: ", err)
		c <- link // the Link value will be resent through the Channel
		return
	}

	fmt.Println(link, " is up!")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	c <- link // the Link value will be resent through the Channel
}
