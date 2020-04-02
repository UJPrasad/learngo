package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardShapes := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Aces", "Two", "Threee", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, shape := range cardShapes {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+shape)
		}
	}

	return cards
}

func (d deck) deal(f int, t int) []string {
	return d[:t]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
