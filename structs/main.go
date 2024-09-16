package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}
type ContactInfo struct {
	email   string
	zipCode int
}

func main() {

	alex := Person{
		firstName: "Alex",
		lastName:  "Paterson",
		ContactInfo: ContactInfo{
			email:   "alex@email.com",
			zipCode: 12345,
		},
	}

	//alexPointer := &alex
	alex.updatName("Nico")
	alex.print()
}

func (pointerToPerson *Person) updatName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
