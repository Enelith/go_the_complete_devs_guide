# Deeper Into Go

## How to Run

To run this program, execute the following command in your terminal:

```bash
go run main.go
```

## Quick overview

The program will have the following methods and functionalities:

- newDeck: Create a list of playing cards. Essentially an array of strings
- print: Log out the contents of a deck of cards
- shuffle: Randomly shuffle the contents of a deck of cards
- deal: Create a "hand" of cards.
- saveToFile: Save a list of cards to a file on the local machine.
- newDeckFromFile: Read and load a list of cards from a file on the local machine.

## Notions

### Arrays & Slices (in a Nutshell)

- an Array is a fixed-length list of elements of the same type.
- a Slice is a dynamically-sized list of elements of the same type. (it can grow and/or shrink)

### For Loops

```bash
for index, card := range cards {
  fmt.Println(card)
}
```
- *index* is the index of the element in the array
- *card* is the current card we're iterating over
- *range cards* takes the slice of 'cards' and loop over it (using the keyword *range*)
- *fmt.Println(card)* run this one time for ,each card in the slice

### Receiver Functions
```bash
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}
```
Any variable of type "deck" now get access to the "print" method
- *d* is the receiver variable. The actualy copy of the deck we're working with is avaiable in the function as a variable called "d". (By convention, name is 1 to 3 letters long related to the type of the receiver)
- *deck* is the type of the receiver variable. Every variable of type 'deck' can call this function on itself.

Note: Go is not a class-based / object-oriented OO language, but it does support methods on types (receiver functions) which vaguely resemble classes in other languages.
Therefore, we never use the terms "class" or "object" in Go. Likewise, Go avoids any mention of "this" or "self".