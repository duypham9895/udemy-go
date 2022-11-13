package main

import (
	"os"
	"testing"
)

func TestNewDesk(t *testing.T) {
	desk := newDesk()

	if len(desk) != 52 {
		t.Errorf("Expected desk length of 52, but got %v", len(desk))
	}

	if desk[0] != "Ace of Spades" {
		t.Errorf("Expected a first card of Ace of Spades, but got %v", desk[0])
	}

	if desk[51] != "King of Clubs" {
		t.Errorf("Expected a last card of King of Clubs, but got %v", desk[0])
	}
}

func TestSaveToFileAndNewDeskFromFile(t *testing.T) {
	const NAME_FILE_TESTING string = "_desk_testing"
	os.Remove(NAME_FILE_TESTING)

	desk := newDesk()
	desk.saveToFile(NAME_FILE_TESTING)

	loadedDesk := newDeskFromFile(NAME_FILE_TESTING)

	if len(desk) != len(loadedDesk) {
		t.Errorf("Expected length of loaded desk equal desk created from code, but they are not equal")
	}

	os.Remove(NAME_FILE_TESTING)
}
