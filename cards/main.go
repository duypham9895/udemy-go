package main


func main() {
	cards := desk{newCard(), "Age of Spades"}
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}