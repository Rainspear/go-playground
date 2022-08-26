package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expect Deck to have 20 elements, but got %d", len(d))
	}
}
