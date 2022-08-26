package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expect Deck to have 20 elements, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expect Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expect Four of Clubs, but got %s", d[len(d)-1])
	}
}
