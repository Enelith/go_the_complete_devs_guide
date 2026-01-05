package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Note: We're going to pretend the methods getGreeting for both englishBot and spanishBot have different body
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
