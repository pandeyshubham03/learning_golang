package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	// card = "Five of Diamonds"

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
