package main

func main() {
	// cards := newDeckFromFile("my_cards.txt")
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.saveToFile("my_cards.txt")
	// fmt.Println(cards.toString())
	// hand, remainCards := deal(cards, 2)
	// hand.print()
	// fmt.Println("--------------------------------")
	// remainCards.print()
}
