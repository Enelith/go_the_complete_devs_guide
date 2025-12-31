package main

import (
	"fmt"
	"runtime"
)

// Definition of what a contactInfo is
type contactInfo struct {
	email   string
	zipcode string
}

// Definition of what a person is
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName:   "Jim",
		lastName:    "Halpert",
		contactInfo: contactInfo{email: "test@test.com", zipcode: "95128"},
	}

	genericPrint(jim)

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jimPointer.print()

	jim.updateNameBis("Jimmy2")
	jim.print()
}

// "*person" is a type description. It means we're working with a pointer to a person.
// "*pointerToPerson" is an operator. It means we want to manipulate the value the pointer is referencing (aka dereferencing)
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *person) updateNameBis(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Note that there's no way to create a generic version of print() from the receiver point of view.
// However, it's doable using params, as follow:
func genericPrint(p interface{}) {
	pc, _, _, _ := runtime.Caller(0)

	// Will display: main.genericPrint: p details
	fmt.Printf("%v: %+v\n", runtime.FuncForPC(pc).Name(), p)
}
