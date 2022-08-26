package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, error := os.ReadFile(filename)
	cardString := ""
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	cardString = string(bs)
	return deck(strings.Split(cardString, ","))
}

func (d deck) shuffle() deck {
	for i := range d {
		newPosition := rand.Intn(len(d) - i)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}
