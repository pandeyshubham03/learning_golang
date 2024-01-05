package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	// card = "Five of Diamonds"
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
