package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

// We see that unlike structs the slice was modified automatically
