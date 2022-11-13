package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	contactAlex := contactInfo{email: "alex@me.com", zip: 700000}
	alex := person{firstName: "Alex", lastName: "Anderson", contact: contactAlex}

	alexPointer := &alex
	fmt.Println(alexPointer)
	alexPointer.updateFirstName("Duy")
	alex.print()
}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
