package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func deal(d desk, handSize int) (desk, desk) {
	return d[:handSize], d[handSize:]
}

func (d desk) toString() string {
	return strings.Join(d, ",")
}

func toByteSlice(value string) []byte {
	return []byte(value)
}

func decodeByteToString(value []byte) string {
	return string(value[:])
}

func (d desk) saveToFile(filename string) error {
	cardsBytes := toByteSlice(d.toString())
	return ioutil.WriteFile(filename, cardsBytes, 0666)
}