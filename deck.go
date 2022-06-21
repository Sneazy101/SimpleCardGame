package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	for _, value := range cardValues {
		for _, suit := range cardsSuits {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handsSize int) (deck, deck) {
	return d[:handsSize], d[handsSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
