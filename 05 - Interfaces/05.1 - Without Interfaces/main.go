package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// In Go, there's no method overloading ~ we leave it like that as for the demonstration,
// but we CLEARLY want to reuse the same code for both types, so how to do from here ?...
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

// Note: We're going to pretend the methods getGreeting for both englishBot and spanishBot have different body
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
