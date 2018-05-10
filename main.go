package main

func main() {
	cards := newDeck()

	// cards.saveToFile("my-cards")
	// fmt.Println(newDeckFromFile("my-cards"))

	cards.shuffle()
	cards.print()
}
