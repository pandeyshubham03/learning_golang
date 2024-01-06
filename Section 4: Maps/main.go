package main

import "fmt"

func main() {

	// method 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// method 2
	// var colors map[string]string

	// method 3
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex Code for", color, "is", hex)
	}
}
