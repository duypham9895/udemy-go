package main

import "fmt"

type desk []string

func newDesk() desk {
	cards := desk{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}

	return cards
}

func (d desk) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}