package main

func main() {
	// var card string = "Ace of Spades"
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// card = "Five of Diamonds"

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
