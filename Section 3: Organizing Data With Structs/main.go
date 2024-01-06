package main

import "fmt"

type person struct {
	// defining a struct
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	var alex person
	fmt.Println(alex)
}
