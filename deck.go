package main

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

func (d deck) print() {
	for _, card := range d {
		println(card)
	}
}
