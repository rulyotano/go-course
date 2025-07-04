package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	fileName := "_decktesting"
	os.Remove(fileName);

	d := newDeck()

	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if !d.Equals(loadedDeck) {
		t.Errorf("Expected loaded deck to be equal to original deck, but they are not")
	}

	os.Remove(fileName);
}