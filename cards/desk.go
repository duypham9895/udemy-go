package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeskFromFile(filename string) desk {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	deskString := decodeByteToString(data)
	cards := strings.Split(deskString, ",")
	return desk(cards)
}

func (d desk) immutableShuffle() desk {
	shuffledCards := make(desk, len(d))

	perm := rand.Perm(len(d))
	for i, randomIdx := range perm {
		shuffledCards[randomIdx] = d[i]
	}

	return shuffledCards
}

func (d desk) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	genRand := rand.New(source)

	genRand.Shuffle(
		len(d),
		func(i, j int) {
			d[i], d[j] = d[j], d[i]
		},
	)
}
