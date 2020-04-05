package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) deal(t int) (deck, deck) {
	return d[:t], d[t:]
}

func (d deck) saveToFile(f string) {

	err := ioutil.WriteFile(f, []byte(d.toString()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func loadFromFile(f string) deck {
	bytes, err := ioutil.ReadFile(f)

	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bytes), ","))
}

func (d deck) shuffle() {
	rand.Seed((time.Now()).UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
