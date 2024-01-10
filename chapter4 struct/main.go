package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "test@test.com", zipCode: 84000,
		},
	}

	//@definition: &variable = give the memory address is pointing at
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	//@definition: *pointer= take the address and give me the value in the address
	(*pointerToPerson).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
