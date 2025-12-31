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
	contact   contactInfo
}

// This definition of person_2 is the same as person, except the field "contact" is implicitly "contactInfo" of type contactInfo
type person_2 struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact:   contactInfo{email: "test@test.com", zipcode: "95128"},
	}

	pam := person_2{
		firstName: "Pam",
		lastName:  "Beesly",
		contactInfo: contactInfo{
			email:   "test@test.com",
			zipcode: "95128",
		},
	}

	jim.print()
	fmt.Printf("Pam: %+v\n", pam)

	genericPrint(jim)
	genericPrint(pam)

	jim.updateName("Jimmy")
	jim.print()
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

// This should update the firstName, but do not (bcs it needs a POINTER)
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
