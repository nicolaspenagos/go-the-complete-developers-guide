package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type Deck []string

// Receiver on a func
// Any variable of type Deck now get
// acess to the print method
// recevier is equivalent of 'this' in JS
// by convention we refer to the receiver
// with one or max two letters name
// that matches the type
func (d Deck) PrintDeck() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func NewDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := fmt.Sprintf("%s of %s", value, suit)
			cards = append(cards, card)
		}
	}

	return cards
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) SaveToFile(filename string) error {
	bytes := []byte(d.ToString())
	return os.WriteFile(filename, bytes, 0666)
}

func NewDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
