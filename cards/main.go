package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println(cards.deal(500, 3))
}
