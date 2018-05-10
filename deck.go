package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

	fmt.Println(d)
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Club", "Diamond"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, error := ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	stringArray := strings.Split(string(byteSlice), ",")

	return deck(stringArray)
}
