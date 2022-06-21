package main

func main() {

	cards := newDeck()
	err := cards.saveToFile("my_cards")
	if err != nil {
		return
	}
}
