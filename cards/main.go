package main

func main() {
	cards := NewDeckFromFile(("my_cards.txt"))
	cards.Shuffle()
	cards.PrintDeck()
}
