package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
