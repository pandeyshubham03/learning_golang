package main

import "fmt"

type contactInfo struct {
	// defining a struct
	email   string
	zipCode int
}

type person struct {
	// defining a struct
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// embedding structs:
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	// jimPointer := &jim
	jim.updateName("Jimmy") // go can directly convert value to pointer automatically
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
