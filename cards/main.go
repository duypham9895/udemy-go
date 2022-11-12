package main


func main() {
	cards := newDesk()

	myCards, _ := deal(cards, 12)

	myCards.shuffle()
	myCards.print()

	// cards.print()

}
