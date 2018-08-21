package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName, lastName string
	contactInfo
}

func main() {
	//adhish := person{"Adhish", "Vishwakarma"}
	// adhish := person{firstName: "Adhish", lastName: "Vishwakarma"}
	// fmt.Println(adhish)
	// var adhish person
	// adhish.firstName = "Adhish"
	// adhish.lastName = "Vishwakarma"
	// fmt.Println(adhish)
	// fmt.Printf("%+v", adhish)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
