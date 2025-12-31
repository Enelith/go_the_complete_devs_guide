package main

import "fmt"

// Definition of what a person is
type person struct {
	firstName string
	lastName  string
}

func main() {
	// Several ways to create a new person (struct).
	alex_1 := person{"Alex", "Smith"}

	fmt.Println("1st way: ", alex_1)

	alex_2 := person{firstName: "Alex", lastName: "Smith"}
	fmt.Println("2nd way: ", alex_2)

	// Note that in this case, the values are set to their zero values (string = "", int & float = 0, bool = false, ...)
	var alex_3 person
	fmt.Println("3rd way: ", alex_3)
	fmt.Printf("3rd way (bis): %+v\n", alex_3)

	// Updating values
	alex_3.firstName = "Alexander"
	alex_3.lastName = "Smith"
	fmt.Printf("3rd way (ter): %+v\n", alex_3)
}
